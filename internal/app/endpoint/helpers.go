package endpoint

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (e *Endpoint) changeGetData(ctx *gin.Context, function func(data []byte) ([]byte, error)) (int, error) {
	writer := ctx.Writer
	body := ctx.Request.Body
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	result, err := function(data)
	if err != nil {
		return http.StatusBadRequest, err
	}

	_, err = writer.Write(result)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return 0, nil
}
