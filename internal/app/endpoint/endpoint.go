package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
)

type Endpoint struct {
	service Service
	engine  *gin.Engine
	log     zerolog.Logger
}

func New(eng *gin.Engine, service Service, out io.Writer) *Endpoint {
	logger := zerolog.New(out)
	logger.With().
		Str("service", "endpoint").
		Logger()
	return &Endpoint{
		engine:  eng,
		service: service,
		log:     logger,
	}
}

func (e *Endpoint) Start(addr string) error {
	logger := e.log.With().
		Str("function", "start").
		Logger()
	logger.Debug().
		Str("status", "start").
		Str("addr", addr)

	logger.Trace().Msg("initialize handlers")
	e.engine.GET("/book", e.BookHandler)
	e.engine.GET("/books", e.BooksHandler)
	e.engine.POST("/book", e.CreateHandler)
	e.engine.PUT("/book", e.UpdateHandler)
	e.engine.DELETE("/book", e.DeleteHandler)

	logger.Trace().
		Str("addr", addr).
		Msg("start listening and serve")
	err := e.engine.Run(addr)
	if err != nil {
		logger.Error().
			Str("status", "failed").
			Err(err)
		return err
	}

	logger.Debug().Str("status", "start")
	return nil
}
