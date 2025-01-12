package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fabiosoliveira/task-api/internal/task"
)

// Lista todas as tarefas.

func CreateTasksHandler(addTask func(name string) task.Task) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Cria uma nova tarefa.
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var newTask task.Task
		err = json.Unmarshal(body, &newTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newTask = addTask(newTask.Name)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	}
}
