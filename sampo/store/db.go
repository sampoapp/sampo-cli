/* This is free and unencumbered software released into the public domain. */

package store

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

type Store struct {
	db *sql.DB
}

// Create
func Create(path string) (*Store, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=rwc", path))
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open failed")
	}
	if _, err := db.Exec("PRAGMA journal_mode=DELETE"); err != nil {
		return nil, errors.Wrap(err, "sql.Exec failed")
	}
	if _, err := db.Exec("PRAGMA encoding=\"UTF-8\""); err != nil {
		return nil, errors.Wrap(err, "sql.Exec failed")
	}
	if _, err := db.Exec("PRAGMA foreign_keys=ON"); err != nil {
		return nil, errors.Wrap(err, "sql.Exec failed")
	}
	if _, err := db.Exec("PRAGMA user_version=1"); err != nil {
		return nil, errors.Wrap(err, "sql.Exec failed")
	}
	return &Store{db}, nil
}

// OpenDefault
func OpenDefault() (*Store, error) {
	return Open("./app.db")
}

// Open
func Open(path string) (*Store, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", path))
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open failed")
	}
	return &Store{db}, nil
}

// Init
func (store *Store) Init(sqlSchema string) error {
	for _, sqlStatement := range strings.Split(sqlSchema, ";\n") {
		if sqlStatement == "" {
			continue // skip empty statements, if any
		}
		if _, err := store.db.Exec(sqlStatement); err != nil {
			return errors.Wrap(err, "sql.Exec failed")
		}
	}
	return nil
}

// CreateUser
func (store *Store) CreateUser(userUUID *uuid.UUID, userNick string, userName string) (int, error) {
	if userUUID == nil {
		newUUID, err := uuid.NewV4()
		if err != nil {
			return 0, errors.Wrap(err, "uuid.NewV4 failed")
		}
		userUUID = &newUUID
	}
	if _, err := store.db.Exec("INSERT INTO user (id, uuid, nick, name) VALUES (NULL, ?, ?, ?)", userUUID.String(), userNick, userName); err != nil {
		return 0, errors.Wrap(err, "sql.Exec failed")
	}
	return 0, nil // TODO
}

// Compact
func (store *Store) Compact() error {
	if _, err := store.db.Exec("VACUUM"); err != nil {
		return errors.Wrap(err, "sql.Exec failed")
	}
	return nil
}

// Close
func (store *Store) Close() error {
	if err := store.db.Close(); err != nil {
		return errors.Wrap(err, "sql.Close failed")
	}
	return nil
}
