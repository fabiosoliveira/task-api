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
