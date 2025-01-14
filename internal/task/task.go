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

func GetTaskById(id int) (Task, error) {
	for _, t := range tasks {
		if id == t.ID {
			return t, nil
		}
	}

	return Task{}, fmt.Errorf("task with ID %d not found", id)
}

func AddTask(name string) Task {
	lastId++
	task := Task{ID: lastId, Name: name}
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int, name string) error {
	for i, t := range tasks {
		if id == t.ID {
			tasks[i].Name = name
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func RemoveTask(id int) error {
	for i, t := range tasks {
		if id == t.ID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}
