package main

import (
	"log"

	"github.com/face4/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer("localhost:8080")
	log.Fatal(srv.ListenAndServe())
}
