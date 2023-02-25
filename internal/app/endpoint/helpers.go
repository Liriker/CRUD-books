package endpoint

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

func (e *Endpoint) changeGetData(ctx *gin.Context, function func(data []byte) ([]byte, error), logger zerolog.Logger) (int, error) {

	logger.Trace().Msg("get body and context")
	writer := ctx.Writer
	body := ctx.Request.Body
	defer body.Close()

	logger.Trace().Msg("read data from body")
	data, err := io.ReadAll(body)
	if err != nil {
		logger.Error().
			Str("status", "failed").
			Err(err)
		return http.StatusInternalServerError, err
	}

	logger.Trace().Msg("get result from function")
	result, err := function(data)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return http.StatusBadRequest, err
	}

	logger.Trace().Msg("write answer to writer")
	_, err = writer.Write(result)
	if err != nil {
		logger.Error().
			Str("status", "failed").
			Err(err)
		return http.StatusInternalServerError, err
	}
	return 0, nil
}
