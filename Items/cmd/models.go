// models.go
// SQL Database Models

package main

import (
	"log"
)

func (m *DataModel) GetAllItems() ([]Item, error) {
	stmt := "SELECT id, title, body, created FROM items;"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err = rows.Scan(&item.Id, &item.Title, &item.Body, &item.Created)
		if err != nil {
			log.Println(err)
		}
		items = append(items, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (m *DataModel) GetOneItem(id int) (*Item, error) {
	stmt := "SELECT id, title, body, created FROM items WHERE id = ?;"
	row := m.DB.QueryRow(stmt, id)

	var item Item
	err := row.Scan(&item.Id, &item.Title, &item.Body, &item.Created)
	if err != nil {
		log.Println(err)
	}

	return &item, nil
}

func (m *DataModel) Insert(title string, body string) error {
	stmt := "INSERT INTO items (title, body) VALUES (?, ?);"
	_, err := m.DB.Exec(stmt, title, body)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (m *DataModel) Update(id int, title string, body string) error {
	stmt := "UPDATE items SET title=?, body=? WHERE id=?;"
	_, err := m.DB.Exec(stmt, title, body, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Delete(id int) error {
	stmt := "DELETE FROM items WHERE id=?;"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
