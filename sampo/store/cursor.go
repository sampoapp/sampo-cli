/* This is free and unencumbered software released into the public domain. */

package store

import (
	"database/sql"
)

type Cursor struct {
	rows *sql.Rows
}

// Close
func (cursor *Cursor) CloseCursor() error {
	return cursor.rows.Close()
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
	return cursor.rows.Scan(dest)
}
