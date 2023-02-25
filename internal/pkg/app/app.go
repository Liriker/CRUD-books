package app

import (
	"CRUD-books/internal/app/endpoint"
	"CRUD-books/internal/app/repository"
	"CRUD-books/internal/app/service"
	"CRUD-books/internal/pkg/app/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
)

func Start(confPath, confFile string) error {
	writer := os.Stdout

	logger := zerolog.New(writer).With().
		Str("service", "app").
		Logger()
	logger.Debug().Str("status", "start")

	logger.Trace().Msg("initialize config")
	conf, err := config.New(confPath, confFile)
	if err != nil {
		return err
	}

	logger.Trace().Msg("initialize repository")
	bd, err := repository.New(conf.User(), conf.Password(), writer)
	if err != nil {
		return err
	}

	logger.Trace().Msg("initialize service")
	serv := service.New(bd, writer)

	logger.Trace().Msg("initialize engine")
	engine := gin.Default()

	logger.Trace().Msg("initialize endpoints")
	ep := endpoint.New(engine, serv, writer)

	logger.Info().Msgf("start host http://%v%v\n", conf.Host(), conf.Port())
	err = ep.Start(conf.Port())
	if err != nil {
		return err
	}

	logger.Debug().Str("status", "done")
	return nil
}
