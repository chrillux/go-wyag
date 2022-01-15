package git

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Object struct {
	obj     ObjectI
	objType string
	repo    *Repository
}

type ObjectI interface {
	Serialize() io.Reader
	Deserialize(data io.Reader)
	GetParents() []string
	String() string
}

func NewObject(repo *Repository, object ObjectI, objType string) *Object {
	o := &Object{
		repo:    repo,
		obj:     object,
		objType: objType,
	}
	return o
}

// Readobject reads a hash and returns the corresponding object.
// A git object structure is an object type, a space, the size as an int, a null byte, and the data.
func ReadObject(hash string) (*Object, error) {
	r := NewRepo()
	objpath := r.RepoFile(filepath.Join(r.Gitdir(), "objects", hash[0:2], hash[2:]), false)
	f, err := os.ReadFile(objpath)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	zread, err := zlib.NewReader(bytes.NewReader(f))
	if err != nil {
		log.Fatal(err)
	}
	buf, err := ioutil.ReadAll(zread)
	if err != nil {
		return nil, err
	}
	// byte 32 is a space
	ispace := bytes.IndexByte(buf, byte(32))
	if ispace < 0 {
		return nil, fmt.Errorf("not valid git object data")
	}
	objType := string(buf[0:ispace])

	// byte 0 is a null byte
	inull := bytes.IndexByte(buf, byte(0))
	if inull < 0 {
		return nil, fmt.Errorf("not valid git object data")
	}
	size, err := strconv.Atoi(string(buf[ispace+1 : inull]))
	if err != nil {
		return nil, err
	}
	if size != len(buf)-inull-1 {
		return nil, fmt.Errorf("malformed object %s: bad length", hash)
	}

	data := bytes.NewReader(buf[inull+1:])
	switch objType {
	case "blob":
		return NewObject(r, NewBlobObject(r, data), "blob"), nil
	case "commit":
		return NewObject(r, NewCommitObject(r, data), "commit"), nil
	}
	return nil, nil
}

// WriteObject by computing the hash, insert header and zlib compress everything. The last part is optional.
func (o *Object) WriteObject(write bool) (*string, error) {
	dataReader := o.Serialize()
	dataBytes, err := ioutil.ReadAll(dataReader)
	if err != nil {
		return nil, err
	}
	result := []byte(strings.Join([]string{o.GetObjType(), " ", fmt.Sprintf("%d", len(dataBytes))}, ""))
	result = append(result, byte(0))
	result = append(result, dataBytes...)
	hash := fmt.Sprintf("%x", sha1.Sum(result))

	if write {
		path := o.repo.RepoFile(filepath.Join("objects", hash[0:2], hash[2:]), true)
		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		err = f.Chmod(os.FileMode(0644))
		if err != nil {
			return nil, err
		}
		// TBD write the data to the file.
		// err = writeData(zlib.NewWriter(f), o.getSerializedData())
		// if err != nil {
		// 	return nil, fmt.Errorf("error writing file: %v", err)
		// }
		// f.Close()
	}
	return &hash, nil
}

func (o *Object) Deserialize(data io.Reader) {
	o.obj.Deserialize(data)
}

func (o *Object) Serialize() io.Reader {
	return o.obj.Serialize()
}

func (o *Object) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *Object) GetObjType() string {
	return o.objType
}

func (o *Object) GetParents() []string {
	return o.obj.GetParents()
}

type KVLM struct {
	KeyValues []KV
	Message   string
}

type KV struct {
	Key   string
	Value string
}

func ParseKeyValueListWithMessage(data io.Reader) *KVLM {
	s := bufio.NewScanner(data)
	curSlice := []string{}
	messageFound := false
	mSlice := []string{}

	kvlm := KVLM{}
	kv := KV{}
	for s.Scan() {
		line := s.Text()
		if messageFound {
			mSlice = append(mSlice, line)
			continue
		}
		if strings.HasPrefix(line, " ") {
			curSlice = append(curSlice, line)
			continue
		} else if !strings.Contains(line, " ") {
			kv.Value = strings.Join(curSlice, "\n")
			kvlm.KeyValues = append(kvlm.KeyValues, kv)
			messageFound = true
			continue
		} else if len(curSlice) > 0 {
			kv.Value = strings.Join(curSlice, "\n")
			kvlm.KeyValues = append(kvlm.KeyValues, kv)
			curSlice = []string{}
		}
		lslice := strings.Split(line, " ")
		if len(lslice) > 1 {
			kv.Key = lslice[0]
			newSlice := []string{}
			for i := 1; i < len(lslice); i++ {
				newSlice = append(newSlice, lslice[i])
			}
			curSlice = append(curSlice, strings.Join(newSlice, " "))
		}
	}
	kvlm.Message = strings.Join(mSlice, "\n")
	return &kvlm
}

func KeyValueListWithMessageSerialize(kvlm KVLM) io.Reader {
	ret := []byte{}
	for _, kv := range kvlm.KeyValues {
		ret = append(ret, []byte(kv.Key)...)
		ret = append(ret, []byte(" ")...)
		ret = append(ret, []byte(strings.ReplaceAll(kv.Value, "\n", "\n "))...)
		ret = append(ret, []byte("\n")...)
	}
	ret = append(ret, []byte("\n")...)
	return bytes.NewReader(ret)
}
