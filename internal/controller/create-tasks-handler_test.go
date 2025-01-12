package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiosoliveira/task-api/internal/task"
)

type MockReader struct{}

func (m *MockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("forced read error")
}

func TestCreateTasksHandler(t *testing.T) {
	mockAddTask := func(name string) task.Task {
		return task.Task{Name: name}
	}

	handler := CreateTasksHandler(mockAddTask)

	t.Run("successful task creation", func(t *testing.T) {
		newTask := task.Task{Name: "Test Task"}
		body, err := json.Marshal(newTask)
		if err != nil {
			t.Errorf("error marshalling request body: %v", err)
		}

		rr := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(body))

		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusCreated)
		}

		var createdTask task.Task
		err = json.NewDecoder(rr.Body).Decode(&createdTask)
		if err != nil {
			t.Errorf("error decoding response body: %v", err)
		}

		if newTask.Name != createdTask.Name {
			t.Errorf("handler returned unexpected body: got %v want %v",
				createdTask.Name, newTask.Name)
		}
	})

	t.Run("invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte("invalid body")))

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusBadRequest)
		}
	})

	t.Run("error reading request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/tasks", &MockReader{})

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusInternalServerError)
		}
	})
}
