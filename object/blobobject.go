package object

import (
	"io"
	"io/ioutil"

	"github.com/chrillux/go-wyag/git"
)

type blobObject struct {
	repo    *git.Repository
	data    io.Reader
	objType string
}

func NewBlobObject(repo *git.Repository, data io.Reader) *blobObject {
	o := &blobObject{
		repo:    repo,
		data:    data,
		objType: "blob",
	}
	return o
}

func (o *blobObject) Deserialize(data io.Reader) {
	o.data = data
}

func (o *blobObject) Serialize() io.Reader {
	return o.data
}

func (o *blobObject) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *blobObject) GetParents() []string {
	return nil
}