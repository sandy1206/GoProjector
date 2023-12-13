package main

import {
	"log"
	"net/http"
	"github.com/gorilla/mux"
}

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

func (s *APIServer) serve() {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// registering the handlers

	log.Println("Starting server on", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subRouter))
}
