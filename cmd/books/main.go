package main

import (
	"log"
	"net/http"
)

const (
	Addr = ":8080"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", getAllBooks)
	mux.HandleFunc("/books/book", getBook)
	mux.HandleFunc("/books/book/update", updateBook)
	mux.HandleFunc("/books/book/delete", deleteBook)
	mux.HandleFunc("/books/create", createBook)

	log.Printf("Run server http://localhost%v", Addr)
	err := http.ListenAndServe(Addr, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Получаем книгу"))
}
func getAllBooks(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Получаем все книги"))
}
func createBook(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Добавляем книгу"))
}
func updateBook(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Изменяем книгу"))
}
func deleteBook(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Удаляем книгу"))
}
