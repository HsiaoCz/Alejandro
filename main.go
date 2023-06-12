package main

import (
	"alejandro/api"
	"alejandro/storage"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello")

	store := storage.NewStorage()
	srv := api.NewServer(store)
	log.Fatal(srv.Start())
}
