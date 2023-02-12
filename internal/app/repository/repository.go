package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func New() (*Repository, error) {
	db, err := sql.Open("mysql", "root:21012001Ilya@/crud-book?parseTime=true")
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Close() error {
	err := r.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) All() (*sql.Rows, error) {

	stmt := `SELECT id, name, author, publish_date FROM book;`

	result, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) Get(id int) *sql.Row {
	stmt := `
			SELECT id, name, author, publish_date FROM book
			WHERE id = ?;
			`

	result := r.db.QueryRow(stmt, id)
	return result
}

func (r *Repository) Create(id int, name, author string, publishDate time.Time) (sql.Result, error) {
	stmt := `
			INSERT INTO book (id, name, author, publish_date)
			VALUES (NULL, ?, ?, ?);
			`

	result, err := r.db.Exec(stmt, id, name, author, publishDate)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) Delete(id int) (sql.Result, error) {
	stmt := `
			DELETE FROM book
			WHERE id = ?;
			`

	result, err := r.db.Exec(stmt, id)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *Repository) Update(row sql.Row) {
	//TODO
}
