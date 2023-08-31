package core

import "fmt"

type Entry struct {
	// SHA-1 Object ID
	oid []byte
	// content
	name string
}

// create a new entry
func (e *Entry) New(oid []byte, name string) error {
	if len(name) == 0 || len(oid) == 0 {
		return fmt.Errorf("name or oid cannot be empty")
	}
	e.oid = oid
	e.name = name
	return nil
}

func (e *Entry) GetName() string {
	return e.name
}
func (e *Entry) GetOid() []byte {
	return e.oid
}
func (e *Entry) SetOid(oid []byte) {
	e.oid = oid
}
