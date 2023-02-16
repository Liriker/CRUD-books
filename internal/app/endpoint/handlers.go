package endpoint

import (
	"github.com/gin-gonic/gin"
	"log"
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
	writer := ctx.Writer

	books, err := e.service.Books()
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = writer.Write(books)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (e *Endpoint) BookHandler(ctx *gin.Context) {
	response, err := e.changeGetData(ctx, e.service.Book)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		log.Println(err)
		return
	}
}

func (e *Endpoint) CreateHandler(ctx *gin.Context) {
	response, err := e.changeGetData(ctx, e.service.CreateBook)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		log.Println(err)
		return
	}
}

func (e *Endpoint) DeleteHandler(ctx *gin.Context) {
	response, err := e.changeGetData(ctx, e.service.DeleteBook)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		log.Println(err)
		return
	}
}

func (e *Endpoint) UpdateHandler(ctx *gin.Context) {
	response, err := e.changeGetData(ctx, e.service.UpdateBook)
	if err != nil {
		ctx.Writer.WriteHeader(response)
		log.Println(err)
		return
	}
}
