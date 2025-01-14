package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiosoliveira/task-api/internal/task"
)

func TestUpdateTaskHandler(t *testing.T) {

	tests := []struct {
		name           string
		task           task.Task
		taskID         string
		updateTaskFunc func(id int, name string) error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:   "successful update",
			task:   task.Task{Name: "Updated Task"},
			taskID: "1",
			updateTaskFunc: func(id int, name string) error {
				return nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:   "invalid task ID",
			task:   task.Task{Name: "Updated Task"},
			taskID: "invalid",
			updateTaskFunc: func(id int, name string) error {
				return nil
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "strconv.Atoi: parsing \"invalid\": invalid syntax\n",
		},
		{
			name:   "task not found",
			task:   task.Task{Name: "Updated Task"},
			taskID: "1",
			updateTaskFunc: func(id int, name string) error {
				return fmt.Errorf("task not found")
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "task with ID 1 not found\n",
		},
		{
			name:   "invalid request body",
			task:   task.Task{},
			taskID: "1",
			updateTaskFunc: func(id int, name string) error {
				return nil
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var req *http.Request
			if tt.name == "invalid request body" {
				req = httptest.NewRequest(http.MethodPut, "/tasks/"+tt.taskID, &MockReader{})
			} else {

				body, _ := json.Marshal(tt.task)

				req = httptest.NewRequest(http.MethodPut, "/tasks/"+tt.taskID, bytes.NewReader(body))
			}
			rr := httptest.NewRecorder()

			mux := &http.ServeMux{}
			mux.HandleFunc("PUT /tasks/{id}", UpdateTaskHandler(tt.updateTaskFunc))
			mux.ServeHTTP(rr, req)

			resp := rr.Result()
			defer resp.Body.Close()

			if tt.expectedStatus != resp.StatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v", resp.StatusCode, tt.expectedStatus)
			}

			if tt.expectedBody != "" {
				respBody, _ := io.ReadAll(resp.Body)
				if tt.expectedBody != string(respBody) {
					t.Errorf("handler returned unexpected body: got %v want %v", string(respBody), tt.expectedBody)
				}
			}
		})
	}
}
