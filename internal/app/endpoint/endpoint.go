package endpoint

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type book interface {
	NewEmpty() *book.Book
	ID() int
	UnmarshalJSON([]byte) error
}

type Endpoint struct {
	books []book
}

func New() *Endpoint {
	return &Endpoint{}
}

func (ep *Endpoint) Books(ctx *gin.Context) {
	writer := ctx.Writer

	req, err := json.Marshal(ep.books)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}

	_, err = writer.Write(req)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
	}
}

func (ep *Endpoint) Book(ctx *gin.Context) {
	writer := ctx.Writer

	body := ctx.Request.Body

}
