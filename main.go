package main

import (
	"alejandro/api"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello")
	srv := api.NewServer("111")
	log.Fatal(srv.Start())
}
