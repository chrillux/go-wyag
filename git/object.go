package git

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"strings"
)

type IGitObject interface {
	serialize(data []byte) error
	deserialize()
	getHash() string
	getSerializedData() []byte
}

type Factory struct {
	obj IGitObject
}

func (f *Factory) serialize(data []byte) error {
	err := f.obj.serialize(data)
	if err != nil {
		return err
	}
	return nil
}

func (f *Factory) deserialize() {
	f.obj.deserialize()
}

func (f *Factory) getHash() string {
	return f.obj.getHash()
}

func (f *Factory) getSerializedData() []byte {
	return f.obj.getSerializedData()
}

type BlobObject struct {
	serializedData []byte
	dataLen        int
	hash           string
}

func (b *BlobObject) serialize(data []byte) error {
	rawdatar := bytes.NewReader(data)
	buf, err := ioutil.ReadAll(rawdatar)
	if err != nil {
		return err
	}
	b.dataLen = len(buf)
	sd := []byte(strings.Join([]string{"blob", fmt.Sprintf("%d", b.dataLen)}, " "))
	sd = append(sd, byte(0))
	sd = append(sd, buf...)
	b.hash = fmt.Sprintf("%x", sha1.Sum(sd))
	b.serializedData = sd
	return nil
}

func (b *BlobObject) deserialize() {
	fmt.Println("commit deserialize")
}

func (b *BlobObject) getHash() string {
	return b.hash
}

func (b *BlobObject) getSerializedData() []byte {
	return b.serializedData
}

func NewObject(objType string) *Factory {
	switch objType {
	case "blob":
		return &Factory{&BlobObject{}}
	}
	return nil
}
