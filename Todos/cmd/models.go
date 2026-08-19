// models.go
// SQL Database Models

package main

import (
	"log"
)

func (m DataModel) GetAllTodos() ([]Todo, error) {
	rows, err := m.DB.Query("SELECT * FROM todos")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Item, &todo.Completed, &todo.Created)
		if err != nil {
			log.Println(err)
		}
		todos = append(todos, todo)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}
	return todos, err
}

func (m DataModel) Insert(item string) error {
	stmt := "INSERT INTO todos (item) VALUES (?)"
	_, err := m.DB.Exec(stmt, item)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m DataModel) Update(id int) error {
	stmt := "UPDATE todos SET completed = 1 WHERE id = ?"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
func (m DataModel) Delete(id int) error {
	stmt := "DELETE FROM todos WHERE id = ?"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
