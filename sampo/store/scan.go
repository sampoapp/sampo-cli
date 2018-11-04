/* This is free and unencumbered software released into the public domain. */

package store

import (
	"database/sql"

	"github.com/sampoapp/sampo-cli/sampo/schema"
)

// ScanEntity
func ScanEntity(rows *sql.Rows) (*schema.Entity, error) {
	var entity schema.Entity
	err := rows.Scan(&entity.ID, &entity.UUID)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
