/* This is free and unencumbered software released into the public domain. */

package schema

// User
type User struct {
	ID   uint64
	UUID string
}

// Entity
type Entity struct {
	ID        uint64
	UUID      string
	CreatedBy uint64
	CreatedAt uint64
	UpdatedBy uint64
	UpdatedAt uint64
}
