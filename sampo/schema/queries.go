/* This is free and unencumbered software released into the public domain. */

package schema

import (
	"github.com/pkg/errors"
	"github.com/sampoapp/sampo-cli/sampo/store"
)

// QueryEntities
func QueryEntities(db *store.Store) (*store.Cursor, error) {
	cursor, err := db.Query("SELECT id, uuid FROM data")
	if err != nil {
		return nil, errors.Wrap(err, "QueryEntities failed")
	}
	return cursor, nil
}
