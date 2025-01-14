package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fabiosoliveira/task-api/internal/task"
)

func GetTaskByIdHandler(getTaskById func(id int) (task.Task, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// recupera uma tarefa.
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t, err := getTaskById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(t)
	}
}
