/* This is free and unencumbered software released into the public domain. */

package store

import "github.com/pkg/errors"

// Query
func (store *Store) Query(sql string, args ...interface{}) (*Cursor, error) {
	rows, err := store.db.Query(sql, args...)
	if err != nil {
		return nil, errors.Wrap(err, "Query failed")
	}
	return &Cursor{rows}, nil
}
