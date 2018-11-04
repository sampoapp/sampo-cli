/* This is free and unencumbered software released into the public domain. */

package store

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Cursor struct {
	rows *sql.Rows
}

// Close
func (cursor *Cursor) CloseCursor() error {
	err := cursor.rows.Close()
	if err != nil {
		return errors.Wrap(err, "CloseCursor failed")
	}
	return nil
}

// Err
func (cursor *Cursor) Err() error {
	return cursor.rows.Err()
}

// Next
func (cursor *Cursor) Next() bool {
	return cursor.rows.Next()
}

// Scan
func (cursor *Cursor) Scan(dest ...interface{}) error {
	err := cursor.rows.Scan(dest...)
	if err != nil {
		return errors.Wrap(err, "Scan failed")
	}
	return nil
}
