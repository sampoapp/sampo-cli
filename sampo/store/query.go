/* This is free and unencumbered software released into the public domain. */

package store

// Query
func (store *Store) Query(sql string) (*Cursor, error) {
	rows, err := store.db.Query(sql)
	if err != nil {
		return nil, err
	}
	return &Cursor{rows}, nil
}
