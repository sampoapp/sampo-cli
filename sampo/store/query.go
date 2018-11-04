/* This is free and unencumbered software released into the public domain. */

package store

import (
	"database/sql"
)

// QueryEntities
func QueryEntities(store *Store) (*sql.Rows, error) {
	rows, err := store.db.Query("SELECT id, uuid FROM data")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
