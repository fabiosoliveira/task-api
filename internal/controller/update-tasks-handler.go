package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/fabiosoliveira/task-api/internal/task"
)

// Atualiza uma tarefa.

func UpdateTaskHandler(updateTask func(id int, name string) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		var task task.Task
		err = json.Unmarshal(body, &task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Atualiza a tarefa com o ID informado.
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		err = updateTask(id, task.Name)
		mu.Unlock()

		if err != nil {
			http.Error(w, fmt.Sprintf("task with ID %d not found", id), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
