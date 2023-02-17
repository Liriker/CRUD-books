package service

import (
	"CRUD-books/internal/app/service/domain/book"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type repository interface {
	All() (*sql.Rows, error)
	Get(id int) *sql.Row
	Create(id int, name, author string, publishDate time.Time) (sql.Result, error)
	Delete(id int) (sql.Result, error)
	Update(id int, name, author string, publishDate time.Time) (sql.Result, error)
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
		b := book.New()

		err = rows.Scan(&b.Id, &b.Name, &b.Author, &b.PublishDate)

		arr = append(arr, b)
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
	var b book.Book

	err := json.Unmarshal(data, &b)
	if err != nil {
		return nil, err
	}

	id := b.ID()

	row := s.db.Get(id)
	err = row.Scan(&b.Id, &b.Name, &b.Author, &b.PublishDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return []byte("There is no book with this ID"), nil
		}
		return nil, err
	}

	body, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *Service) CreateBook(data []byte) ([]byte, error) {
	var b book.Book

	err := json.Unmarshal(data, &b)
	if err != nil {
		return nil, err
	}

	_, err = s.db.Create(b.Id, b.Name, b.Author, b.PublishDate)

	body := fmt.Sprintf("Book #%v was created\n", b.ID())

	return []byte(body), nil
}

func (s *Service) DeleteBook(data []byte) ([]byte, error) {
	var b book.Book

	err := json.Unmarshal(data, &b)
	if err != nil {
		return nil, err
	}

	id := b.ID()

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
	var newBook book.Book
	var b book.Book

	err := json.Unmarshal(data, &newBook)
	if err != nil {
		return nil, err
	}

	id := newBook.ID()

	row := s.db.Get(id)
	err = row.Scan(&b.Id, &b.Name, &b.Author, &b.PublishDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return []byte("There is no book with this ID"), nil
		}
		return nil, err
	}
	b.UpdateBook(newBook)

	_, err = s.db.Update(b.Id, b.Name, b.Author, b.PublishDate)
	if err != nil {
		return nil, err
	}

	body := fmt.Sprintf("Book #%v was updated", id)

	return []byte(body), nil
}
