package model

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Setup() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	var err error
	db, err = sql.Open("sqlite3", "todos.sqlite")
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping db: %w", err)
	}

	return db, nil
}

func MakeMigrations() error {
	db, err := Setup()
	if err != nil {
		return fmt.Errorf("failed to setup db: %w", err)
	}

	q := `CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            todo VARCHAR(128) NULL,
            done boolean
         );`

	_, err = db.Exec(q)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
