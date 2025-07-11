# 📡 WebSocket Task Board API (Go)

A real-time task board API built in **Go (Golang)** using **WebSockets** and **REST**. Supports multiple clients connected simultaneously to collaborate on tasks with live updates, powered by **Gorilla WebSocket** and a clean architecture.

---

## 🌟 Features

### 🔄 Real-Time WebSocket API
- Live **task creation**, **updates**, and **deletion**
- Broadcasts events to all connected clients

### 🌐 RESTful API
- `POST /tasks` → Create new tasks
- `GET /tasks` → Get all tasks
- `PUT /tasks/{id}` → Update a task
- `DELETE /tasks/{id}` → Delete a task

### 🧠 In-Memory Task Store
- Fast, thread-safe task storage using `sync.RWMutex`
- Can be extended to Redis/PostgreSQL later

### 🔧 Built with Go
- Lightweight and concurrent by design
- Uses:
  - `github.com/gorilla/websocket`
  - `github.com/gorilla/mux`

---

## 📁 Project Structure

```

websocket-task-board/
├── cmd/
│   └── main.go                # Main entry point
├── internal/
│   └── board/
│       ├── handler.go         # WebSocket logic
│       ├── manager.go         # Client manager
│       ├── store.go           # In-memory task store
│       ├── types.go           # Shared types (Task, Message)
│       └── http.go            # REST handlers
├── go.mod

````

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/JawherKl/api-design.git
cd websocket/golang/websocket-task-board
go mod tidy
````

### 2. Run the Server

```bash
go run cmd/main.go
```

* Server runs at: `http://localhost:8080`
* WebSocket endpoint: `ws://localhost:8080/ws`

---

## 📡 WebSocket Protocol

### Message Format

```json
{
  "type": "task_created",
  "payload": {
    "id": "uuid",
    "title": "Task title",
    "status": "todo"
  }
}
```

### Supported Message Types

| Type           | Description                |
| -------------- | -------------------------- |
| `task_created` | Broadcast on new task      |
| `task_updated` | Broadcast on task update   |
| `task_deleted` | Broadcast on task deletion |

---

## 🧪 Example Usage (via `wscat`)

```bash
npx wscat -c ws://localhost:8080/ws
```

Send a task:

```json
{
  "type": "task_created",
  "payload": {
    "title": "Learn Go WebSocket",
    "status": "todo"
  }
}
```

---

## 🌍 REST API Reference

| Method | Endpoint      | Description       |
| ------ | ------------- | ----------------- |
| GET    | `/tasks`      | Get all tasks     |
| POST   | `/tasks`      | Create a new task |
| PUT    | `/tasks/{id}` | Update a task     |
| DELETE | `/tasks/{id}` | Delete a task     |

### Example:

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Write docs", "status":"in_progress"}'
```

---

## 🛠️ Tech Stack

| Layer       | Tool                 |
| ----------- | -------------------- |
| Language    | Go (1.20+)           |
| HTTP Router | Gorilla Mux          |
| WebSocket   | Gorilla WebSocket    |
| REST API    | `net/http`, mux      |
| Storage     | In-Memory (sync.Map) |

---

## 📌 Future Improvements

* [ ] Persist tasks with Redis/PostgreSQL
* [ ] Add frontend interface (HTML + JS or React)
* [ ] Add authentication (JWT or Keycloak)
* [ ] Add unit and integration tests
* [ ] Containerize with Docker

---

## 🤝 Contributing

Contributions are welcome! Feel free to:

* Open issues or bugs
* Submit pull requests
* Suggest improvements

---

## 📄 License

MIT License

---

## ✨ Author

**Jawher Kallel**
[GitHub](https://github.com/JawherKl) — [LinkedIn](https://www.linkedin.com/in/jawher-kallel/)
