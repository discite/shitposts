package model

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Setup() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", "todos.sqlite")
	if err != nil {
		fmt.Println("Could not connect to db", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping db", err)
	}
	return db
}

func MakeMigrations() error {
	db := Setup()
	q := `CREATE TABLE IF NOT EXISTS todos (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            todo VARCHAR(128) NULL,
            done boolean
         );`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
