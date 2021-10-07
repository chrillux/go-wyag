package git

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type IGitObject interface {
	serialize(data []byte) error
	deserialize()
	getHash() string
	getSerializedData() []byte
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

func NewObject(objType string) IGitObject {
	switch objType {
	case "blob":
		return &BlobObject{}
	}
	return nil
}

func (r *gitRepository) HashObject(objType string, data []byte, write bool) (*string, error) {
	o := NewObject(objType)
	err := o.serialize(data)
	if err != nil {
		return nil, fmt.Errorf("problem serializing object: %v", err)
	}

	// switch ob := o.(type) {
	// case *BlobObject:
	// 	fmt.Println("it's a blob!!", ob)
	// }
	hash := o.getHash()

	if write {
		path := r.repoFile(filepath.Join("objects", hash[0:2], hash[2:]), true)
		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		err = f.Chmod(os.FileMode(0644))
		if err != nil {
			return nil, err
		}
		err = writeData(zlib.NewWriter(f), o.getSerializedData())
		if err != nil {
			return nil, fmt.Errorf("error writing file: %v", err)
		}
		f.Close()
	}
	return &hash, nil
}

func writeData(f io.WriteCloser, dataToWrite []byte) error {
	defer f.Close()
	_, err := f.Write(dataToWrite)
	if err != nil {
		return err
	}
	return nil
}
