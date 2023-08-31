package core

import "os"

const (
	// metadata directory
	GMetadataDir = ".gitgo"
	/**
	Chmod 0744 (chmod a+rwx,g-wx,o-wx,ug-s,-t) sets permissions so that, (U)ser / owner can read, can write and can execute. (G)roup can read, can't write and can't execute. (O)thers can read, can't write and can't execute.
	*/
	GDefaultPermissions = os.FileMode(0744)
	GMetadataDirContent = "objects|refs"
)
