package task

import "fmt"

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks = []Task{}

var lastId = 0

func GetTasks() []Task {
	return tasks
}

func findTaskById(id int) int {
	for i, t := range tasks {
		if id == t.ID {
			return i
		}
	}

	return -1
}

func GetTaskById(id int) (*Task, error) {

	index := findTaskById(id)
	if index != -1 {
		return &tasks[index], nil
	}

	return &Task{}, fmt.Errorf("task with ID %d not found", id)
}

func AddTask(name string) *Task {
	lastId++
	task := Task{ID: lastId, Name: name}
	tasks = append(tasks, task)
	return &task
}

func UpdateTask(id int, name string) error {
	index := findTaskById(id)
	if index != -1 {
		tasks[index].Name = name
		return nil
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func RemoveTask(id int) error {
	index := findTaskById(id)
	if index != -1 {
		tasks = append(tasks[:index], tasks[index+1:]...)
		return nil
	}

	return fmt.Errorf("task with ID %d not found", id)
}
