// init_db.go
// Initialize Database Connection Pool

package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// initDB handles connection setup and returns the active connection pool.
func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/go_web_apps?parseTime=true")
	if err != nil {
		return nil, err
	}

	// Verify the credentials and connection are actually valid.
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
