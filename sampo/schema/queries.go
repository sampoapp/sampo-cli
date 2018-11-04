/* This is free and unencumbered software released into the public domain. */

package schema

import (
	"database/sql"

	"github.com/sampoapp/sampo-cli/sampo/store"
)

// QueryEntities
func QueryEntities(db *store.Store) (*sql.Rows, error) {
	rows, err := db.Query("SELECT id, uuid FROM data")
	if err != nil {
		return nil, err
	}
	return rows, nil
}
