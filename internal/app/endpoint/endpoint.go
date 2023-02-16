package endpoint

import "github.com/gin-gonic/gin"

type Endpoint struct {
	service Service
	engine  *gin.Engine
}

func New(eng *gin.Engine, service Service) *Endpoint {
	return &Endpoint{
		engine:  eng,
		service: service,
	}
}

func (e *Endpoint) Start() error {
	e.engine.GET("/book", e.BookHandler)
	e.engine.GET("/books", e.BooksHandler)
	e.engine.POST("/book", e.CreateHandler)
	e.engine.PUT("/book", e.UpdateHandler)
	e.engine.DELETE("/book", e.DeleteHandler)

	err := e.engine.Run()
	if err != nil {
		return err
	}

	return nil
}
