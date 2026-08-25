// models.go
// SQL Database Models

package main

import (
	"log"
)

func (m *DataModel) GetAllMovies() ([]Movie, error) {
	stmt := `SELECT id, title, year, runtime, genre, created FROM movies ORDER BY created DESC LIMIT 5;`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Runtime, &movie.Genre, &movie.Created)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return movies, nil
}

func (m *DataModel) GetOneMovie(id int) (Movie, error) {
	stmt := "SELECT id, title, year, runtime, genre, created FROM movies WHERE id = ?;"
	row := m.DB.QueryRow(stmt, id)

	var movie Movie
	err := row.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Runtime, &movie.Genre, &movie.Created)
	if err != nil {
		log.Println(err)
	}

	return movie, nil
}

func (m *DataModel) Insert(title string, year int, runtime int, genre string) (int, error) {
	stmt := "INSERT INTO movies(title, year, runtime, genre) VALUES(?, ?, ?, ?);"
	result, err := m.DB.Exec(stmt, title, year, runtime, genre)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *DataModel) Update(id int, title string, year int, runtime int, genre string) error {
	stmt := "UPDATE movies SET title=?, year=?, runtime=?, genre=? WHERE id=?;"
	_, err := m.DB.Exec(stmt, title, year, runtime, genre, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Delete(id int) error {
	stmt := "DELETE FROM movies WHERE id=?;"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
