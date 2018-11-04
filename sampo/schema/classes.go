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

// TODO
type Account struct {
	Entity
}

// A record of the address of a web page serving as a shortcut to it.
type Bookmark struct {
	Entity
}

// TODO
type Contract struct {
	Entity
}

// TODO
type Document struct {
	Entity
}

// TODO
type Event struct {
	Entity
}

// TODO
type Network struct {
	Entity
}

// TODO
type Note struct {
	Entity
}

// TODO
type Paper struct {
	Entity
}

// TODO
type Payment struct {
	Entity
}

// TODO
type Person struct {
	Entity
}

// TODO
type Place struct {
	Entity
}

// TODO
type Project struct {
	Entity
}

// TODO
type Quote struct {
	Entity
}

// TODO
type Task struct {
	Entity
}
