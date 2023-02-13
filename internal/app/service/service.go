package service

import (
	"CRUD-books/internal/app/service/domain/book"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"
)

type repository interface {
	All() (*sql.Rows, error)
	Get(id int) *sql.Row
	Create(id int, name, author string, publishDate time.Time) (sql.Result, error)
	Delete(id int) (sql.Result, error)
	Update()
}

type Service struct {
	db repository
}

func New(rep repository) *Service {
	return &Service{
		db: rep,
	}
}

func (s *Service) Books() ([]byte, error) {
	var arr []*book.Book

	rows, err := s.db.All()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := book.New()

		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.PublishDate)

		arr = append(arr, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(arr)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Book(data []byte) ([]byte, error) {
	book := *book.New()

	err := json.Unmarshal(data, &book)
	if err != nil {
		return nil, err
	}

	id := book.ID()

	row := s.db.Get(id)
	err = row.Scan(&book.Id, &book.Name, &book.Author, &book.PublishDate)
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(book)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *Service) CreateBook(data []byte) ([]byte, error) {
	book := book.New()

	err := json.Unmarshal(data, book)
	if err != nil {
		return nil, err
	}

	date := time.Unix(book.PublishDate, 0)

	result, err := s.db.Create(book.Id, book.Name, book.Author, date)

	insertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	body := fmt.Sprintf("Book #%v was created\n", insertId)

	return []byte(body), nil
}

func (s *Service) DeleteBook(data []byte) ([]byte, error) {
	book := book.New()

	err := json.Unmarshal(data, book)
	if err != nil {
		return nil, err
	}

	id := book.ID()

	result, err := s.db.Delete(id)
	if err != nil {
		return nil, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	body := fmt.Sprintf("Book #%v was deleted", insertId)

	return []byte(body), nil
}

func (s *Service) UpdateBook(data []byte) ([]byte, error) {
	book := book.New()

	err := json.Unmarshal(data, book)
	if err != nil {
		return nil, err
	}

	id := book.ID()

	// TODO update book

	body := fmt.Sprintf("Book #%v was updated", id)

	return []byte(body), nil
}
