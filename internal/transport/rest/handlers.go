package rest

import (
	"CRUD-books/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", booksHandler)
	mux.HandleFunc("/books/book", bookHandler)

	return mux
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id := 1
		book, err := getBook(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		_, err = w.Write([]byte(book))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case http.MethodPost:
		id := 11
		operation, err := updateBook(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		_, err = w.Write([]byte(operation))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case http.MethodDelete:
		id := 11
		operation, err := deleteBook(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		_, err = w.Write([]byte(operation))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		operation, err := getAllBooks()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		_, err = w.Write([]byte(operation))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		book := domain.EmptyBook()
		err = json.Unmarshal(body, book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		action, err := createBook(*book)

		w.Write([]byte(action))
	}
}

func getBook(id int) (string, error) {
	return fmt.Sprintf("Полуачем книгу по id = %v", id), nil
}

func getAllBooks() (string, error) {
	return fmt.Sprintf("Получаем все книги"), nil
}
func createBook(book domain.Book) (string, error) {
	return fmt.Sprintf("Создали книгу:\n%v", book), nil
}
func updateBook(id int) (string, error) {
	return fmt.Sprintf("Обновляем книгу id = %v", id), nil
}
func deleteBook(id int) (string, error) {
	return fmt.Sprintf("Удаляем книгу id = %v", id), nil
}
