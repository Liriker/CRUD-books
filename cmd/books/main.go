package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	r.GET("/book", booksHandler)
	r.GET("/book", bookHandler)
	r.POST("/book", createHandler)
	r.PUT("/book", updateHandler)
	r.DELETE("/book", deleteHandler)

	log.Println("Start server to http://localhost:8080")

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
