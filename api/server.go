package api

import (
	"alejandro/conf"
	"alejandro/storage"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type server struct {
	ua *userApi
}

func NewServer(store *storage.Storage) *server {
	return &server{
		ua: NewUserApi(store),
	}
}

func (s *server) Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", s.handleGetUserName).Methods("GET")
	srv := http.Server{
		Handler:      r,
		Addr:        fmt.Sprintf("%s:%s", conf.Conf.AC,conf.Conf.AC),
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	return srv.ListenAndServe()
}
