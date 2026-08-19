// models.go
// Database Models

package main

import (
	"log"
)

func (m *DataModel) GetAllRecords() ([]Record, error) {
	stmt := "SELECT * FROM students;"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.Id, &record.Name, &record.Age, &record.Major, &record.Created)
		if err != nil {
			log.Println(err)
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}
	return records, err
}

func (m *DataModel) GetOneRecord(id int) (*Record, error) {
	stmt := "SELECT id, name, age, major, created FROM students WHERE id = ?;"
	row := m.DB.QueryRow(stmt, id)
	var record Record
	err := row.Scan(&record.Id, &record.Name, &record.Age, &record.Major, &record.Created)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (m *DataModel) Insert(name string, age string, major string) error {
	stmt := "INSERT INTO students (name, age, major) VALUES (?, ?, ?)"
	_, err := m.DB.Exec(stmt, name, age, major)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Update(id int, name string, age string, major string) error {
	stmt := "UPDATE students SET name = ?, age = ?, major = ? WHERE id = ?"
	_, err := m.DB.Exec(stmt, name, age, major, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
func (m *DataModel) Delete(id int) error {
	stmt := "DELETE FROM students WHERE id = ?;"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
