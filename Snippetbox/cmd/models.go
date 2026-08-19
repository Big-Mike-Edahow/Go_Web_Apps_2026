// models.go
// SQL Database Models

package main

import (
	"log"
)

func (m *DataModel) Insert(title string, content string) error {
	stmt := "INSERT INTO snippets(title, content) VALUES(?, ?);"
	_, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (m *DataModel) Update(id int, title string, content string) error {
	stmt := "UPDATE snippets SET title=?, content=? WHERE id=?;"
	_, err := m.DB.Exec(stmt, title, content, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Delete(id int) error {
	stmt := "DELETE FROM snippets WHERE id=?;"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) GetOneSnippet(id int) (*Snippet, error) {
	stmt := "SELECT id, title, content, created FROM snippets WHERE id = ?;"
	row := m.DB.QueryRow(stmt, id)

	var snippet Snippet
	err := row.Scan(&snippet.Id, &snippet.Title, &snippet.Content, &snippet.Created)
	if err != nil {
		log.Println(err)
	}

	return &snippet, nil
}

func (m *DataModel) GetAllSnippets() ([]Snippet, error) {
	stmt := "SELECT id, title, content, created FROM snippets ORDER BY id DESC LIMIT 5;"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var snippets []Snippet
	for rows.Next() {
		var snippet Snippet
		err = rows.Scan(&snippet.Id, &snippet.Title, &snippet.Content, &snippet.Created)
		if err != nil {
			log.Println(err)
		}
		snippets = append(snippets, snippet)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
