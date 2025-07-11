package board

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"` // e.g., "todo", "in_progress", "done"
}

type Message struct {
	Type    string      `json:"type"`    // e.g., "task_created", "task_updated"
	Payload interface{} `json:"payload"` // can be a Task or other structs
}
