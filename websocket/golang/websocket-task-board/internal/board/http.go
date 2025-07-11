package board

import (
	"encoding/json"
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := taskStore.All()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
