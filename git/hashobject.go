package git

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func (r *gitRepository) HashObject(objType string, data []byte, write bool) (*string, error) {
	o := NewObject(objType)
	err := o.serialize(data)
	if err != nil {
		return nil, fmt.Errorf("problem serializing object: %v", err)
	}

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
