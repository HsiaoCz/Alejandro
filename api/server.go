package api

import (
	"alejandro/conf"
	"alejandro/storage"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

type server struct {
	ua *userApi
}

func NewServer(store *storage.Storage) *server {
	return &server{
		ua: NewUserApi(store),
	}
}

func (s *server) Start() error {
	r.GET("/user/:name", s.ua.registerRouter)
	srv := http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", conf.Conf.AC.Addr, conf.Conf.AC.Port),
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	return srv.ListenAndServe()
}
