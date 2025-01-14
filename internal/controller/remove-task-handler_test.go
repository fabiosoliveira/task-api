package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoveTaskHandler(t *testing.T) {
	tests := []struct {
		name           string
		taskID         string
		removeTaskFunc func(id int) error
		expectedStatus int
	}{
		{
			name:   "Valid task ID",
			taskID: "1",
			removeTaskFunc: func(id int) error {
				return nil
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:   "Invalid task ID",
			taskID: "abc",
			removeTaskFunc: func(id int) error {
				return nil
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:   "Task not found",
			taskID: "2",
			removeTaskFunc: func(id int) error {
				return errors.New("task not found")
			},
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/tasks/"+tt.taskID, nil)
			rr := httptest.NewRecorder()

			mux := &http.ServeMux{}
			mux.HandleFunc("DELETE /tasks/{id}", RemoveTaskHandler(tt.removeTaskFunc))
			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}
		})
	}
}
