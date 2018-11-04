/* This is free and unencumbered software released into the public domain. */

package schema

import (
	"database/sql"
)

// ScanEntity
func ScanEntity(rows *sql.Rows) (*Entity, error) {
	var entity Entity
	err := rows.Scan(&entity.ID, &entity.UUID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
