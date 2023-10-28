package model

import (
	"errors"
)

type Todo struct {
	Id   uint64 `json:"id"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}

func (t *Todo) Create() error {
	db := Setup()
	q := `INSERT INTO todos (Todo, Done) VALUES(?, ?)`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(t.Todo, t.Done)
	if err != nil {
		return err
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	return nil
}

func (t *Todo) Get() (Todo, error) {
	db := Setup()
	q := `SELECT Id, Todo, Done FROM todos WHERE Id=?`
	todo := Todo{}
	err := db.QueryRow(q, t.Id).Scan(&todo.Id, &todo.Todo, &todo.Done)
	return todo, err
}

func (t *Todo) GetAll() ([]Todo, error) {
	db := Setup()
	q := `SELECT Id, Todo, Done FROM todos`
	rows, err := db.Query(q)
	if err != nil {
		return []Todo{}, err
	}
	defer rows.Close()
	todos := []Todo{}
	for rows.Next() {
		rows.Scan(
			&t.Id,
			&t.Todo,
			&t.Done,
		)
		todos = append(todos, *t)
	}
	return todos, nil
}

func (t *Todo) Update() error {
	db := Setup()
	q := `UPDATE todos set Todo=?, Done=? WHERE id=?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(t.Todo, t.Done, t.Id)
	if err != nil {
		return err
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	return nil
}

func (t *Todo) MarkDone() error {
	todo, err := t.Get()
	if err != nil {
		return err
	}
	todo.Done = !todo.Done
	todo.Update()
	if err != nil {
		return err
	}
	return err
}

func (t *Todo) Delete(id uint64) error {
	db := Setup()
	q := `DELETE FROM todos WHERE id=?`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	r, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: Se esperaba una fila afectada")
	}
	return nil
}
