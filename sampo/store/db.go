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

type Record = map[string]interface{}

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
func (store *Store) CreateUser(userUUID *uuid.UUID, userNick string, userName string) (int64, error) {
	if userUUID == nil {
		newUUID, err := uuid.NewV4()
		if err != nil {
			return 0, errors.Wrap(err, "uuid.NewV4 failed")
		}
		userUUID = &newUUID
	}
	result, err := store.db.Exec("INSERT INTO user (id, uuid, nick, name) VALUES (NULL, ?, ?, ?)", userUUID.String(), userNick, userName)
	if err != nil {
		return 0, errors.Wrap(err, "sql.Exec failed")
	}
	return result.LastInsertId()
}

// CreateEntity
func (store *Store) CreateEntity(entity Record) (int64, error) {
	if _, ok := entity["uuid"]; !ok {
		entityUUID, err := uuid.NewV4()
		if err != nil {
			return 0, errors.Wrap(err, "uuid.NewV4 failed")
		}
		entity["uuid"] = entityUUID.String()
	}
	result, err := store.db.Exec("INSERT INTO data (id, uuid, created_by, created_at, updated_by, updated_at) VALUES (NULL, ?, ?, ?, ?, ?)", entity["uuid"], 1, 0, nil, nil) // TODO
	if err != nil {
		return 0, errors.Wrap(err, "sql.Exec failed")
	}
	return result.LastInsertId()
}

func canonicalizeValue(anyValue interface{}) interface{} {
	switch value := anyValue.(type) {
	case nil:
		return nil
	case int:
		return value
	case string:
		return strings.TrimSpace(value)
	default:
		return nil // TODO
	}
}

// CreateEntityOfClass
func (store *Store) CreateEntityOfClass(className string, entityID int64, entity Record) (int64, error) {
	var sqlColumns, sqlVariables []string
	var sqlBindings []interface{}
	for column, value := range entity {
		sqlColumns = append(sqlColumns, column)
		sqlVariables = append(sqlVariables, "?")
		sqlBindings = append(sqlBindings, canonicalizeValue(value))
	}
	sqlCommand := fmt.Sprintf("INSERT INTO data_%s (%s) VALUES (%s)", className, strings.Join(sqlColumns, ", "), strings.Join(sqlVariables, ", "))
	//fmt.Println(sqlCommand, sqlBindings) // DEBUG
	if _, err := store.db.Exec(sqlCommand, sqlBindings...); err != nil {
		return 0, errors.Wrap(err, "sql.Exec failed")
	}
	return entityID, nil
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
