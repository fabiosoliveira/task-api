package controller

import (
	"net/http"
	"strconv"
)

func RemoveTaskHandler(removeTask func(id int) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Remove uma tarefa.
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mu.Lock()
		err = removeTask(id)
		mu.Unlock()

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
