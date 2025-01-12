package task

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
