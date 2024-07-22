package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {

	ms := &mockStore{}
	service := NewTaskService()

	t.Run("should return an error if name is empty", func(t *testing.T) {
		payload := &Task{
			Name: "",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.handleCreateTask)
	})
}
