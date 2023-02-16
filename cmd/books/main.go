package main

import (
	"CRUD-books/internal/pkg/app"
	"log"
)

func main() {
	err := app.Start()
	if err != nil {
		log.Println(err)
	}
}
