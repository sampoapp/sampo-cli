/* This is free and unencumbered software released into the public domain. */

package store

import (
	"database/sql"
)

// Query
func (store *Store) Query(sql string) (*sql.Rows, error) {
	rows, err := store.db.Query(sql)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
