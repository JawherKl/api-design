package board

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := taskStore.All()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	task.ID = uuid.NewString()
	taskStore.Add(task)

	manager.Broadcast <- Message{
		Type:    "task_created",
		Payload: task,
	}

	json.NewEncoder(w).Encode(task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Read error", http.StatusBadRequest)
		return
	}

	var task Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	task.ID = id // force ID
	if ok := taskStore.Update(task); !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	manager.Broadcast <- Message{
		Type:    "task_updated",
		Payload: task,
	}

	json.NewEncoder(w).Encode(task)
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if ok := taskStore.Delete(id); !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	manager.Broadcast <- Message{
		Type:    "task_deleted",
		Payload: map[string]string{"id": id},
	}

	w.WriteHeader(http.StatusNoContent)
}
