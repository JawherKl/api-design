package board

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	clientID := uuid.NewString()
	client := &Client{
		ID:   clientID,
		Conn: conn,
		Send: make(chan Message),
	}

	manager.Register <- client

	// Start goroutines
	go readMessages(client)
	go writeMessages(client)
}

func readMessages(client *Client) {
	defer func() {
		manager.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, msgBytes, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var msg Message
		if err := json.Unmarshal(msgBytes, &msg); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}

		log.Printf("Message from %s: %+v\n", client.ID, msg)

		switch msg.Type {
			case "task_created":
				taskData := msg.Payload.(map[string]interface{})
				task := Task{
					ID:     uuid.NewString(),
					Title:  taskData["title"].(string),
					Status: taskData["status"].(string),
				}
				taskStore.Add(task)

				manager.Broadcast <- Message{
					Type:    "task_created",
					Payload: task,
				}

			case "task_updated":
				taskData := msg.Payload.(map[string]interface{})
				task := Task{
					ID:     taskData["id"].(string),
					Title:  taskData["title"].(string),
					Status: taskData["status"].(string),
				}
				if taskStore.Update(task) {
					manager.Broadcast <- Message{
						Type:    "task_updated",
						Payload: task,
					}
				}

			case "task_deleted":
				taskData := msg.Payload.(map[string]interface{})
				id := taskData["id"].(string)
				if taskStore.Delete(id) {
					manager.Broadcast <- Message{
						Type:    "task_deleted",
						Payload: map[string]string{"id": id},
					}
				}
		}
	}
}

func writeMessages(client *Client) {
	for msg := range client.Send {
		client.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := client.Conn.WriteJSON(msg); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
