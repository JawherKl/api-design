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

    srv := &http.Server{
        Addr:    ":8080",
        Handler: r,
    }

    log.Println("Server running on http://localhost:8080")
    log.Fatal(srv.ListenAndServe())
}
