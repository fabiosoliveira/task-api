package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fabiosoliveira/task-api/internal/task"
)

// Lista todas as tarefas.

func TasksHandler(getTasks func() []task.Task) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		mu.RLock()
		task := getTasks()
		mu.RUnlock()

		response, err := json.Marshal(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, string(response))
	}
}
