package main

import (
	"net/http"

	"github.com/fabiosoliveira/task-api/internal/controller"
	"github.com/fabiosoliveira/task-api/internal/task"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", controller.TasksHandler(task.GetTasks))
	mux.HandleFunc("GET /tasks/{id}", controller.GetTaskByIdHandler(task.GetTaskById))
	mux.HandleFunc("POST /tasks", controller.CreateTasksHandler(task.AddTask))
	mux.HandleFunc("PUT /tasks/{id}", controller.UpdateTaskHandler(task.UpdateTask))
	mux.HandleFunc("DELETE /tasks/{id}", controller.RemoveTaskHandler(task.RemoveTask))

	http.ListenAndServe(":8080", mux)
}
