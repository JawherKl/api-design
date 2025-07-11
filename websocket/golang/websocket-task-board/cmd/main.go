package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/JawherKl/websocket-task-board/internal/board"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/ws", board.HandleWebSocket)
	r.HandleFunc("/tasks", board.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", board.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", board.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/tasks/{id}", board.DeleteTaskHandler).Methods("DELETE")

    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    log.Println("Server running on http://localhost:8080")
    log.Fatal(srv.ListenAndServe())
}
