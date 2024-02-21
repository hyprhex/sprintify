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
	memoryStorage := &MockStore{}
	service := NewTaskService(memoryStorage)

	t.Run("Should return an error if name is empty", func(t *testing.T) {
		payload := &Task{
			Name: "",
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.handleCreateTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Error("invlaid status code, it should fail")
		}
	})

	t.Run("should create a task", func(t *testing.T) {
		payload := &Task{
			Name:         "Createing a unit test",
			ProjectID:    1,
			AssignedToID: 42,
		}

		b, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks", service.handleCreateTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("invlaid status code %d, got %d", http.StatusCreated, rr.Code)
		}
	})
}

func TestGetTask(t *testing.T) {
	memoryStorage := &MockStore{}
	service := NewTaskService(memoryStorage)

	t.Run("should return the task", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/tasks/42", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/tasks/{id}", service.handleGetTask)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Error("invlaid status code, it should fail")
		}
	})
}
