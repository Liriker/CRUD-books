package app

import (
	"CRUD-books/internal/app/endpoint"
	"CRUD-books/internal/app/repository"
	"CRUD-books/internal/app/service"
	"CRUD-books/internal/pkg/app/config"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
)

func Start(confName, confPath string, output io.Writer) error {

	logger := zerolog.New(output).With().
		Str("service", "app").
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	logger.Trace().Msg("initialize config")
	conf, err := config.New(confName, confPath)
	if err != nil {
		return err
	}

	logger.Trace().Msg("initialize repository")
	bd, err := repository.New(conf.User(), conf.Password(), output)
	logger.Info().
		Str("user", conf.User()).
		Str("pass", conf.Password()).
		Msg("")
	if err != nil {
		return err
	}

	logger.Trace().Msg("initialize service")
	serv := service.New(bd, output)

	logger.Trace().Msg("initialize engine")
	engine := gin.Default()

	logger.Trace().Msg("initialize endpoints")
	ep := endpoint.New(engine, serv, output)

	logger.Info().Msgf("start host http://%v%v\n", conf.Host(), conf.Port())
	err = ep.Start(conf.Port())
	if err != nil {
		return err
	}

	logger.Debug().Str("status", "done").Msg("")
	return nil
}
