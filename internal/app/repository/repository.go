package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"io"
	"strconv"
	"time"
)

type Repository struct {
	db  *sql.DB
	log zerolog.Logger
}

func New(user, password string, logOut io.Writer) (*Repository, error) {
	dataSourceName := fmt.Sprintf("%v:%v@/crud-books?parseTime=true", user, password)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	logger := zerolog.New(logOut).With().
		Str("service", "repository").
		Logger()

	return &Repository{
		db:  db,
		log: logger,
	}, nil
}

func (r *Repository) Close() error {
	logger := r.log.With().Str("function", "close").Logger()

	logger.Debug().Str("status", "start").Msg("")

	err := r.db.Close()
	if err != nil {
		logger.Debug().
			Str("status", "failed").
			Err(err).Msg("")
		return err
	}

	logger.Debug().Str("status", "done")
	return nil
}

func (r *Repository) All() (*sql.Rows, error) {
	logger := r.log.With().Str("function", "all").Logger()
	logger.Debug().Str("status", "start").Msg("")

	logger.Trace().Msg("initialize query")
	stmt := `SELECT id, name, author, publish_date FROM books;`

	logger.Trace().Msg("get rows")
	result, err := r.db.Query(stmt)
	if err != nil {
		logger.Warn().
			Str("status", "failed").
			Err(err).Msg("")
		return nil, err
	}

	logger.Debug().Str("status", "done").Msg("")
	return result, nil
}

func (r *Repository) Get(id int) *sql.Row {
	logger := r.log.With().
		Str("function", "Get").
		Str("id", strconv.Itoa(id)).
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	logger.Trace().Msg("initialize query")
	stmt := `
			SELECT id, name, author, publish_date FROM books
			WHERE id = ?;
			`

	logger.Trace().Msg("get row")
	result := r.db.QueryRow(stmt, id)

	logger.Debug().Str("status", "done").Msg("")
	return result
}

func (r *Repository) Create(id int, name, author string, publishDate time.Time) (sql.Result, error) {
	logger := r.log.With().
		Str("function", "Create").
		Str("id", strconv.Itoa(id)).
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	logger.Trace().Msg("initialize query")
	stmt := `
			INSERT INTO books (id, name, author, publish_date)
			VALUES (?, ?, ?, ?);
			`

	logger.Trace().Msg("initialize null-ID")
	nullId := sql.NullInt32{
		Int32: int32(id),
		Valid: true,
	}
	logger.Trace().Msg("check id to null")
	if id <= 0 {
		nullId.Valid = false
	}

	logger.Trace().Msg("get result")
	result, err := r.db.Exec(stmt, nullId, name, author, publishDate)
	if err != nil {
		return nil, err
	}

	logger.Debug().Str("status", "done").Msg("")
	return result, nil
}

func (r *Repository) Delete(id int) (sql.Result, error) {
	logger := r.log.With().
		Str("function", "delete").
		Str("id", strconv.Itoa(id)).
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	logger.Trace().Msg("initialize query")
	stmt := `
			DELETE FROM books
			WHERE id = ?;
			`

	logger.Trace().Msg("get result")
	result, err := r.db.Exec(stmt, id)
	if err != nil {
		logger.Warn().
			Str("status", "fail").
			Err(err).Msg("")
		return nil, err
	}

	logger.Debug().Str("status", "done").Msg("")
	return result, nil
}

func (r *Repository) Update(id int, name, author string, publishDate time.Time) (sql.Result, error) {
	logger := r.log.With().
		Str("function", "update").
		Str("id", strconv.Itoa(id)).
		Logger()
	logger.Debug().Str("status", "start").Msg("")

	logger.Trace().Msg("initialize query")
	stmt := `
			UPDATE books
			SET name = ?,
				author = ?,
				publish_date = ?
			WHERE id = ?;
			`

	logger.Trace().Msg("get result")
	result, err := r.db.Exec(stmt, name, author, publishDate, id)
	if err != nil {
		logger.Warn().
			Str("status", "fail").
			Err(err).Msg("")
		return nil, err
	}

	logger.Debug().Str("status", "done").Msg("")
	return result, nil
}
