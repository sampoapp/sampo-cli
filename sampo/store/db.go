/* This is free and unencumbered software released into the public domain. */

package store

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Store struct {
	db *sql.DB
}

// OpenDefault
func OpenDefault() (*Store, error) {
	return Open("./app.db")
}

// Open
func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "Open failed")
	}
	return &Store{db}, nil
}

// Close
func (store *Store) Close() error {
	err := store.db.Close()
	if err != nil {
		return errors.Wrap(err, "Close failed")
	}
	return nil
}
