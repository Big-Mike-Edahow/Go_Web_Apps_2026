// models.go
// SQL Database Models

package main

import (
	"log"
)

func (m *DataModel) GetAllBooks() ([]Book, error) {
	stmt := "SELECT * FROM bookstore;"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Price, &book.Excerpt, &book.Created)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (m *DataModel) GetOneBook(id int) (*Book, error) {
	stmt := "SELECT id, isbn, title, author, price, excerpt, created FROM bookstore WHERE id = ?;"
	row := m.DB.QueryRow(stmt, id)

	var book Book
	err := row.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Price, &book.Excerpt, &book.Created)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (m *DataModel) Insert(isbn string, title string, author string, price float32, excerpt string) error {
	stmt := "INSERT INTO bookstore (isbn, title, author, price, excerpt) VALUES (?, ?, ?, ?, ?);"
	_, err := m.DB.Exec(stmt, isbn, title, author, price, excerpt)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (m *DataModel) Update(id int, isbn string, title string, author string, price float32, excerpt string) error {
	stmt := "UPDATE bookstore SET isbn=?, title=?, author=?, price=?, excerpt=? WHERE id=?;"
	_, err := m.DB.Exec(stmt, isbn, title, author, price, excerpt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Delete(id int) error {
	stmt := "DELETE FROM bookstore WHERE id=?;"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
