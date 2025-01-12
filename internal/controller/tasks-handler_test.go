package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiosoliveira/task-api/internal/task"
)

func TestTasksHandler(t *testing.T) {
	mockTasks := []task.Task{
		{ID: 1, Name: "Task 1"},
		{ID: 2, Name: "Task 2"},
	}

	getTasks := func() []task.Task {
		return mockTasks
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)

	handler := TasksHandler(getTasks)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			rr.Header().Get("Content-Type"), "application/json")
	}

	expectedResponse, err := json.Marshal(mockTasks)
	if err != nil {
		t.Errorf("error marshalling expected response: %v", err)
	}

	if rr.Body.String() != string(expectedResponse) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expectedResponse))
	}
}
