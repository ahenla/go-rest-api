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

func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// registering services
	usersService := NewUserService(s.store)
	usersService.RegisterRoutes(subrouter)

	projectService := NewProjectService(s.store)
	projectService.RegisterRoutes(subrouter)

	tasksService := NewTaskService(s.store)
	tasksService.RegisterRoutes(subrouter)

	log.Println("Starting server on port", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
