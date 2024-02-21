package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

// Initialize the router, register all the services
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userService := NewUserService(s.store)
	userService.RegisterRouter(subRouter)

	taskService := NewTaskService(s.store)
	taskService.RegisterRoutes(subRouter)

	log.Println("Starting the API server at", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subRouter))
}
