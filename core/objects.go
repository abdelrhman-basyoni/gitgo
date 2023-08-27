package core

type ObjectInterface interface {
	ToString() string
	GetOid() []byte
	SetOid([]byte)
}
