// models.go
// SQL Database Models

package main

import (
	"log"
)

func (m *DataModel) GetAllPosts() ([]Post, error) {
    stmt := "SELECT * FROM posts;"
	rows, err := m.DB.Query(stmt)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.Created)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}
	return posts, err
}

func (m *DataModel) GetOnePost(id int) (*Post, error) {
    stmt := "SELECT id, title, content, created FROM posts WHERE id = ?;"
	row := m.DB.QueryRow(stmt, id)
	var post Post
	err := row.Scan(&post.Id, &post.Title, &post.Content, &post.Created)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (m *DataModel) Insert(title string, content string) error {
	stmt := "INSERT INTO posts (title, content) VALUES (?, ?);"
	_, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Update(id int, title string, content string) error {
	stmt := "UPDATE posts SET title = ?, content = ? WHERE id = ?;"
	_, err := m.DB.Exec(stmt, title, content, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) Delete(id int) error {
	stmt := "DELETE FROM comments WHERE post_id = ?;"
	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	
	stmt = "DELETE FROM posts WHERE id = ?;"
	_, err = m.DB.Exec(stmt, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (m *DataModel) GetComments(id int) ([]Comment, error) {
    stmt := "SELECT id, content, created FROM comments WHERE post_id = ? ORDER BY id DESC;"
	rows, err := m.DB.Query(stmt, id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.Id, &comment.Content, &comment.Created)
		if err != nil {
			log.Println(err)
		}
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
	}
	return comments, err
}

func (m *DataModel) InsertComment(postId int, content string) error {
    stmt := "INSERT INTO comments (post_id, content) VALUES (?, ?);"
	_, err := m.DB.Exec(stmt, postId, content)
	if err != nil {
		log.Println(err)
	}
	return err
}
