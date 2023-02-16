package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func New(user, password string) (*Repository, error) {
	dataSourceName := fmt.Sprintf("%v:%v@/crud-books?parseTime=true", user, password)
	db, err := sql.Open("mysql", dataSourceName)
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

	stmt := `SELECT id, name, author, publish_date FROM books;`

	result, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Repository) Get(id int) *sql.Row {
	stmt := `
			SELECT id, name, author, publish_date FROM books
			WHERE id = ?;
			`

	result := r.db.QueryRow(stmt, id)
	return result
}

func (r *Repository) Create(id int, name, author string, publishDate time.Time) (sql.Result, error) {
	stmt := `
			INSERT INTO books (id, name, author, publish_date)
			VALUES (?, ?, ?, ?);
			`

	nullId := sql.NullInt32{
		Int32: int32(id),
		Valid: true,
	}
	if id <= 0 {
		nullId.Valid = false
	}

	result, err := r.db.Exec(stmt, nullId, name, author, publishDate)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *Repository) Delete(id int) (sql.Result, error) {
	stmt := `
			DELETE FROM books
			WHERE id = ?;
			`

	result, err := r.db.Exec(stmt, id)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (r *Repository) Update() {
	//TODO update Row
}
