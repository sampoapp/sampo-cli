/* This is free and unencumbered software released into the public domain. */

package schema

import (
	"github.com/pkg/errors"
	"github.com/sampoapp/sampo-cli/sampo/store"
)

// ScanEntity
func ScanEntity(cursor *store.Cursor) (*Entity, error) {
	var entity Entity
	err := cursor.Scan(&entity.ID, &entity.UUID)
	if err != nil {
		return nil, errors.Wrap(err, "ScanEntity failed")
	}
	return &entity, nil
}
