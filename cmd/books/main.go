package main

import (
	"CRUD-books/internal/pkg/app"
	"github.com/rs/zerolog"
	"os"
)

const (
	ConfPath = "./configs"
	ConfFile = "app"
)

func main() {
	logger := zerolog.New(os.Stdout)

	err := app.Start(ConfPath, ConfFile)
	if err != nil {
		logger.Fatal().Err(err)
	}
}
