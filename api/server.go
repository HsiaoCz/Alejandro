package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	listenAddr string
}

func NewServer(listenAddr string) *server {
	return &server{
		listenAddr: listenAddr,
	}
}

func (s *server) Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", s.handleGetUserName).Methods("GET")
	srv := http.Server{
		Handler:      r,
		Addr:         s.listenAddr,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	return srv.ListenAndServe()
}
