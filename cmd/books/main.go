package main

import (
	"CRUD-books/internal/pkg/app"
	"log"
)

const (
	ConfPath = "./configs"
	ConfFile = "app"
)

func main() {
	err := app.Start(ConfPath, ConfFile)
	if err != nil {
		log.Println(err)
	}
}
