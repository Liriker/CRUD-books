package main

import (
	"CRUD-books/internal/pkg/app"
	"CRUD-books/pkg/logs"
	"github.com/rs/zerolog"
)

const (
	ConfPath = "./configs"
	ConfFile = "app"
)

func main() {
	logs.ParseLogLvl()

	output := logs.Writer()
	logger := zerolog.New(output)

	err := app.Start(ConfFile, ConfPath, *output)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
}
