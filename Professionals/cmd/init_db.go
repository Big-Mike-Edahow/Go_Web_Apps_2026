// init_db.go
// Initialize Database Connection

package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/go_web_apps?parseTime=true")
	if err != nil {
		return nil, err
	}

	// Verify the credentials and connection are actually valid.
	if err := db.Ping(); err != nil {
		db.Close() // Clean up if the ping fails.
		return nil, err
	}

	return db, nil
}
