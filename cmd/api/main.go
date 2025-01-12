package main

import (
	"net/http"

	"github.com/fabiosoliveira/task-api/internal/controller"
	"github.com/fabiosoliveira/task-api/internal/task"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", controller.TasksHandler(task.GetTasks))
	mux.HandleFunc("POST /tasks", controller.CreateTasksHandler(task.AddTask))

	http.ListenAndServe(":8080", mux)
}
