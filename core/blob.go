package core

type Blob struct {
	// SHA-1 Object ID
	oid []byte
	// content
	data []byte
}

func (b *Blob) New(data []byte) error {
	b.data = data

	return nil
}

func (b *Blob) GetOid() []byte {
	return b.oid
}
func (b *Blob) SetOid(oid []byte) {
	b.oid = oid
}
func (b *Blob) ToString() string {
	return string(b.data)
}
