package app

import (
	"CRUD-books/internal/app/endpoint"
	"CRUD-books/internal/app/repository"
	"CRUD-books/internal/app/service"
	"CRUD-books/internal/pkg/app/config"
	"github.com/gin-gonic/gin"
	"log"
)

func Start() error {
	conf := config.New()

	bd, err := repository.New(conf.User, conf.Password)
	if err != nil {
		return err
	}

	serv := service.New(bd)

	engine := gin.Default()

	ep := endpoint.New(engine, serv)

	err = ep.Start()
	if err != nil {
		return err
	}
	log.Println("start host http://localhost:8080")

	return nil
}
