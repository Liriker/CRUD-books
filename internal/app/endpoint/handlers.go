package endpoint

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service interface {
	Books() ([]byte, error)
	Book(data []byte) ([]byte, error)
	CreateBook(data []byte) ([]byte, error)
	DeleteBook(data []byte) ([]byte, error)
	UpdateBook(data []byte) ([]byte, error)
}

func (e *Endpoint) BooksHandler(ctx *gin.Context) {
	logger := e.log.With().
		Str("handler", "books").
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	logger.Trace().Msg("get writer")
	writer := ctx.Writer

	logger.Trace().Msg("get books array")
	books, err := e.service.Books()
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err).Msg("")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Trace().Msg("write books array")
	_, err = writer.Write(books)
	if err != nil {
		logger.Error().
			Str("status", "failed").
			Err(err).Msg("")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (e *Endpoint) BookHandler(ctx *gin.Context) {
	logger := e.log.With().
		Str("handler", "book").
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	response, err := e.changeGetData(ctx, e.service.Book, logger)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		return
	}

	logger.Debug().Str("status", "done").Msg("")
}

func (e *Endpoint) CreateHandler(ctx *gin.Context) {
	logger := e.log.With().
		Str("handler", "create").
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	response, err := e.changeGetData(ctx, e.service.CreateBook, logger)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		return
	}

	logger.Debug().Str("status", "done").Msg("")
}

func (e *Endpoint) DeleteHandler(ctx *gin.Context) {
	logger := e.log.With().
		Str("handler", "delete").
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	response, err := e.changeGetData(ctx, e.service.DeleteBook, logger)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		return
	}

	logger.Debug().Str("status", "done").Msg("")
}

func (e *Endpoint) UpdateHandler(ctx *gin.Context) {
	logger := e.log.With().
		Str("handler", "update").
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	response, err := e.changeGetData(ctx, e.service.UpdateBook, logger)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		return
	}

	logger.Debug().Str("status", "done").Msg("")
}
