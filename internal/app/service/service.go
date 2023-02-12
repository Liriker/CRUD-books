package service

import (
	"CRUD-books/internal/app/service/domain/book"
	"database/sql"
)

type repository interface {
	All() (*sql.Rows, error)
}

type Service struct {
	db repository
}

func New(rep repository) *Service {
	return &Service{
		db: rep,
	}
}

func (s *Service) Books() ([]*book.Book, error) {
	var arr []*book.Book

	rows, err := s.db.All()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := book.NewEmpty()

		err = rows.Scan(&book.Id, &book.Name, &book.Author, &book.PublishDate)

		arr = append(arr, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return arr, nil
}
