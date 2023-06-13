package main

import (
	"alejandro/api"
	"alejandro/conf"
	"alejandro/storage"
	"log"
)

func main() {
	if err := conf.InitConf(); err != nil {
		log.Fatal(err)
	}
	store := storage.NewStorage()
	srv := api.NewServer(store)
	log.Fatal(srv.Start())
}
