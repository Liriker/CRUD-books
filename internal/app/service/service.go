package service

import (
	"CRUD-books/internal/app/service/domain/book"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
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
	db  repository
	log zerolog.Logger
}

func New(rep repository, out io.Writer) *Service {
	logger := zerolog.New(out).
		With().
		Str("service", "service").
		Logger()

	return &Service{
		db:  rep,
		log: logger,
	}
}

func (s *Service) Books() ([]byte, error) {
	logger := s.log.With().
		Str("function", "books").
		Logger()
	logger.Debug().Str("status", "start")

	var arr []*book.Book

	logger.Trace().Msg("get rows")
	rows, err := s.db.All()
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}
	defer rows.Close()

	logger.Trace().Msg("put rows to array")
	for rows.Next() {
		b := book.New()

		err = rows.Scan(&b.Id, &b.Name, &b.Author, &b.PublishDate)

		arr = append(arr, b)
	}

	err = rows.Err()
	logger.Trace().Msg("check rows to errors")
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("marshal array to json")
	result, err := json.Marshal(arr)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Debug().Str("status", "done")
	return result, nil
}

func (s *Service) Book(data []byte) ([]byte, error) {
	logger := s.log.With().
		Str("function", "book").
		Logger()
	logger.Debug().Str("status", "start")

	var b book.Book

	logger.Trace().Msg("unmarshal json")
	err := json.Unmarshal(data, &b)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("get id")
	id := b.ID()

	logger.Trace().Msg("get row")
	row := s.db.Get(id)
	logger.Trace().Msg("import row value to book")
	err = row.Scan(&b.Id, &b.Name, &b.Author, &b.PublishDate)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		if err == sql.ErrNoRows {
			return []byte("There is no book with that ID"), nil
		}
		return nil, err
	}

	logger.Trace().Msg("marshal book to json")
	body, err := json.Marshal(b)
	if err != nil {
		logger.Warn().
			Str("status", "filed").
			Err(err)
		return nil, err
	}

	logger.Debug().Str("status", "done")
	return body, nil
}

func (s *Service) CreateBook(data []byte) ([]byte, error) {
	logger := s.log.With().
		Str("function", "createBook").
		Logger()
	logger.Debug().Str("status", "start")

	var b book.Book

	logger.Trace().Msg("unmarshal json data")
	err := json.Unmarshal(data, &b)
	if err != nil {
		log.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("create book in repo")
	_, err = s.db.Create(b.Id, b.Name, b.Author, b.PublishDate)

	logger.Trace().Msg("form answer")
	body := fmt.Sprintf("Book #%v was created\n", b.ID())

	logger.Debug().Str("status", "done")
	return []byte(body), nil
}

func (s *Service) DeleteBook(data []byte) ([]byte, error) {
	logger := s.log.With().
		Str("function", "deleteBook").
		Logger()
	logger.Debug().Str("status", "start")

	var b book.Book

	logger.Trace().Msg("unmarshal json data")
	err := json.Unmarshal(data, &b)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("get id to var")
	id := b.ID()

	logger.Trace().Msg("delete book in repo")
	result, err := s.db.Delete(id)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("get last insert id")
	insertId, err := result.LastInsertId()
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("form answer")
	body := fmt.Sprintf("Book #%v was deleted", insertId)

	logger.Debug().Str("status", "done")
	return []byte(body), nil
}

func (s *Service) UpdateBook(data []byte) ([]byte, error) {
	logger := s.log.With().
		Str("function", "updateBook").
		Logger()
	logger.Debug().Str("status", "start")

	var newBook book.Book
	var b book.Book

	logger.Trace().Msg("unmarshal json data")
	err := json.Unmarshal(data, &newBook)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("get book id")
	id := newBook.ID()

	logger.Trace().Msg("get present data")
	row := s.db.Get(id)
	err = row.Scan(&b.Id, &b.Name, &b.Author, &b.PublishDate)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		if err == sql.ErrNoRows {
			return []byte("There is no book with this ID"), nil
		}
		return nil, err
	}

	logger.Trace().Msg("update book entity to new data")
	b.UpdateBook(newBook)

	logger.Trace().Msg("update book in repo")
	_, err = s.db.Update(b.Id, b.Name, b.Author, b.PublishDate)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err)
		return nil, err
	}

	logger.Trace().Msg("form answer")
	body := fmt.Sprintf("Book #%v was updated", id)

	logger.Debug().Str("status", "done")
	return []byte(body), nil
}
