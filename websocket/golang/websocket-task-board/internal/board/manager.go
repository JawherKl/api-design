package board

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan Message
}

type Manager struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	mu         sync.Mutex
}

var manager = NewManager()

func NewManager() *Manager {
	m := &Manager{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}

	go m.start()
	return m
}

func (m *Manager) start() {
	for {
		select {
		case client := <-m.Register:
			m.mu.Lock()
			m.Clients[client.ID] = client
			m.mu.Unlock()
			log.Println("Client registered:", client.ID)

		case client := <-m.Unregister:
			m.mu.Lock()
			if _, ok := m.Clients[client.ID]; ok {
				delete(m.Clients, client.ID)
				close(client.Send)
				log.Println("Client unregistered:", client.ID)
			}
			m.mu.Unlock()

		case msg := <-m.Broadcast:
			m.mu.Lock()
			for _, client := range m.Clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(m.Clients, client.ID)
				}
			}
			m.mu.Unlock()
		}
	}
}
