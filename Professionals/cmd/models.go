// models.go
// Database Models

package main

import (
	"log"
)

func (m *DataModel) GetAllRecords() ([]Record, error) {
	rows, err := m.DB.Query("SELECT * FROM profs;")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var records []Record
	for rows.Next() {
		var record Record
		err := rows.Scan(&record.Id, &record.Name, &record.Age, &record.Employ, &record.Created)
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
	row := m.DB.QueryRow("SELECT id, name, age, employ, created FROM profs WHERE id = ?;", id)
	var record Record
	err := row.Scan(&record.Id, &record.Name, &record.Age, &record.Employ, &record.Created)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (m *DataModel) Insert(name string, age string, employ string) error {
	stmt := "INSERT INTO profs (name, age, employ) VALUES (?, ?, ?)"
	_, err := m.DB.Exec(stmt, name, age, employ)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Update(id int, name string, age string, employ string) error {
	stmt := "UPDATE profs SET name = ?, age = ?, employ = ? WHERE id = ?"
	_, err := m.DB.Exec(stmt, name, age, employ, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Delete(id int) error {
	stmt := "DELETE FROM profs WHERE id = ?;"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
