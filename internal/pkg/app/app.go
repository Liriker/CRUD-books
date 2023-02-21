package app

import (
	"CRUD-books/internal/app/endpoint"
	"CRUD-books/internal/app/repository"
	"CRUD-books/internal/app/service"
	"CRUD-books/internal/pkg/app/config"
	"github.com/gin-gonic/gin"
	"log"
)

func Start(confPath, confFile string) error {
	conf, err := config.New(confPath, confFile)
	if err != nil {
		return err
	}

	bd, err := repository.New(conf.UserAndPassword())
	if err != nil {
		return err
	}

	serv := service.New(bd)

	engine := gin.Default()

	ep := endpoint.New(engine, serv)

	log.Printf("start host http://%v%v\n", conf.Host(), conf.Port())
	err = ep.Start(conf.Port())
	if err != nil {
		return err
	}

	return nil
}
