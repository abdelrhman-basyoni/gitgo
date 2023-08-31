package core

import (
	"fmt"
	"sort"
)

type Tree struct {
	// SHA-1 Object ID
	oid []byte
	// Entries
	data []Entry
}

func (t *Tree) New(entries []Entry) error {
	t.data = entries
	return nil
}

func (t *Tree) GetOid() []byte {
	return t.oid
}

func (t *Tree) SetOid(oid []byte) {
	t.oid = oid
}

// Return a serialized representation of the tree
func (t *Tree) ToString() string {
	dataCopy := make([]Entry, len(t.data))
	copy(dataCopy, t.data)
	// Sort the entries by name
	sort.SliceStable(dataCopy, func(i, j int) bool {
		return dataCopy[i].GetName() < dataCopy[j].GetName()
	})
	// pack all entires into []byte as follows:
	// <Type> <size>\x00<mode> <name>\x00<oid><mode> <name>\x00<oid>...

	var entries []byte

	for _, entry := range dataCopy {
		//Append  the mode , name and null-byte
		// Example: 100644 .gitingore\x00<oid>   ( the first 2 digits for the file type, next 4 digits for file mode )
		entries = append(entries, []byte(fmt.Sprintf("10%04o %s%x", GDefaultPermissions, entry.GetName(), 0x00))...)
		//Append the oid
		entries = append(entries, entry.GetOid()...)
	}
	return string(entries)
}
