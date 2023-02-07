package main

import (
	"CRUD-books/internal/transport/rest"
	"log"
	"net/http"
)

const (
	Addr = ":8080"
)

func main() {
	mux := rest.GetMux()

	log.Printf("Run server http://localhost%v", Addr)
	err := http.ListenAndServe(Addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}
