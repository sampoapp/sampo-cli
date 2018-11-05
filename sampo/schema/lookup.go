/* This is free and unencumbered software released into the public domain. */

package schema

import (
	"github.com/pkg/errors"
	"github.com/sampoapp/sampo-cli/sampo/store"
)

// LookupEntity
func LookupEntity(db *store.Store, uuid string) (*Entity, error) {
	var entity *Entity

	cursor, err := db.Query("SELECT id, uuid FROM data WHERE uuid = ?", uuid)
	if err != nil {
		return nil, errors.Wrap(err, "LookupEntity failed")
	}
	defer cursor.CloseCursor()

	for cursor.Next() {
		entity, err = ScanEntity(cursor)
		if err != nil {
			return nil, errors.Wrap(err, "LookupEntity failed")
		}
	}

	err = cursor.Err()
	if err != nil {
		return nil, errors.Wrap(err, "LookupEntity failed")
	}
	return entity, nil
}
