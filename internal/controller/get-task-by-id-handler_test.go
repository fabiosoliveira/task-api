package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/fabiosoliveira/task-api/internal/task"
)

func TestGetTaskByIdHandler(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockTask       task.Task
		mockError      error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "valid id",
			id:             "1",
			mockTask:       task.Task{ID: 1, Name: "Test Task"},
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"Test Task"}`,
		},
		{
			name:           "invalid id",
			id:             "abc",
			mockTask:       task.Task{},
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "strconv.Atoi: parsing \"abc\": invalid syntax",
		},
		{
			name:           "task not found",
			id:             "2",
			mockTask:       task.Task{},
			mockError:      errors.New("task not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   "task not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/tasks/"+tt.id, nil)

			rr := httptest.NewRecorder()

			getTaskById := func(id int) (task.Task, error) {
				if strconv.Itoa(id) == tt.id {
					return tt.mockTask, tt.mockError
				}
				return task.Task{}, errors.New("task not found")
			}

			mux := &http.ServeMux{}
			mux.HandleFunc("GET /tasks/{id}", GetTaskByIdHandler(getTaskById))
			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			if body := strings.TrimSpace(rr.Body.String()); body != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v",
					body, tt.expectedBody)
			}
		})
	}
}
