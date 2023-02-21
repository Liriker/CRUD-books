package db

import "github.com/kelseyhightower/envconfig"

type DB struct {
	user     string
	password string
}

func New() (*DB, error) {
	var db DB

	if err := envconfig.Process("db", &db); err != nil {
		return nil, err
	}

	return &db, nil
}

func (db *DB) User() string {
	return db.user
}

func (db *DB) Password() string {
	return db.password
}
