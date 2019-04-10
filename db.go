package coresql

import (
	"database/sql"

	"github.com/xo/dburl"
)

// DB represents a wrapper for SQL DB providing extra methods.
type DB struct {
	*sql.DB
}

// Open will attempt to open a new database connection.
func Open(u string) (*sql.DB, error) {
	db, err := dburl.Open(u)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Check will attempt to ping the database to see if the connection is still alive.
func (db *DB) Check() ([]string, bool) {
	if err := db.Ping(); err != nil {
		return []string{err.Error()}, false
	}
	return []string{"database connection ok"}, true
}
