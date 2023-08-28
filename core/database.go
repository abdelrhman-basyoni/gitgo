package core

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type databaseInterface interface {
	// Init the Database
	New(path string) error
	//store new object to the database
	Store(object ObjectInterface) error
}

type Database struct {
	path string
}

func (db *Database) New(path string) error {
	db.path = path
	fmt.Printf("new Database %v", path)

	return nil
}

func (db *Database) Store(object ObjectInterface) error {
	// <type> <size>\x00 <content>
	//Example: blob 5\x00hello

	var encodedData bytes.Buffer

	data := []byte(fmt.Sprintf("%T %d\x00", object, len(object.ToString())))

	data = append(data, object.ToString()...)

	//Encode the Data
	if err := binary.Write(&encodedData, binary.BigEndian, data); err != nil {
		return err
	}

	// Create a SHA-1 hash of the encoded data
	objectId := sha1.Sum(encodedData.Bytes())
	object.SetOid(objectId[:])

	return db.writeObject(object.GetOid(), encodedData.Bytes())
}

func (db *Database) writeObject(oid []byte, data []byte) error {
	oidHex := fmt.Sprintf("%x", oid)
	// .git/objects/<first 2 characters from OId>/ <rest of Oid>
	objectDir := strings.Join([]string{db.path, oidHex[:2]}, string(os.PathSeparator))
	objectPath := strings.Join([]string{objectDir, oidHex[2:]}, string(os.PathSeparator))

	// Create the ObjectPath and all its parents
	if err := os.MkdirAll(objectDir, GDefaultPermissions); err != nil {
		return err
	}

	// Get a temporary directory matches the os
	tmpDir := os.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, db.generateTempObjectName(oidHex))
	if err != nil {
		return err
	}
	defer tmpFile.Close()

	//compress the data with zlib deflate
	var compressedData bytes.Buffer
	zw := zlib.NewWriter(&compressedData)
	defer zw.Close()
	if _, err := zw.Write(data); err != nil {
		return err
	}

	// write data to the temp file
	if _, err = tmpFile.Write(compressedData.Bytes()); err != nil {
		return err
	}

	// move the temp file to the object path
	return os.Rename(tmpFile.Name(), objectPath)
}

func (db *Database) generateTempObjectName(oidHex string) string {
	return fmt.Sprintf("tmp_Obj_%x", oidHex[:3])
}
