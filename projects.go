package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type ProjectService struct {
	store Store
}

func NewProjectService(s Store) *ProjectService {
	return &ProjectService{store: s}
}

func (s *ProjectService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/projects", WithJWTAuth(s.handleCreateProject, s.store)).Methods("POST")
	r.HandleFunc("/projects/{id}", WithJWTAuth(s.handleSingleProject, s.store)).Methods("GET", "DELETE")
}

func (s *ProjectService) handleCreateProject(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload!"})
		return
	}

	defer r.Body.Close()

	var project *Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invlaid request payload!"})
		return
	}

	err = validateProjectPayload(project)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	p, err := s.store.CreateProject(project)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating project"})
		return
	}

	WriteJSON(w, http.StatusCreated, p)
}

func (s *ProjectService) handleSingleProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "id is required"})
		return
	}

	if r.Method == "GET" {
		p, err := s.store.GetProject(id)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Project not found"})
			return
		}

		WriteJSON(w, http.StatusOK, p)

	}

	if r.Method == "DELETE" {
		err := s.store.DeleteProject(id)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Project not found"})
			return
		}

		WriteJSON(w, http.StatusOK, map[string]string{"message": "Project deleted"})

	}
}

func validateProjectPayload(project *Project) error {
	if project.Name == "" {
		return errNameRequired
	}

	return nil
}
