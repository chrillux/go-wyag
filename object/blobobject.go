package object

import (
	"io"
	"io/ioutil"

	"github.com/chrillux/go-wyag/git"
)

type blobObject struct {
	repo *git.Repository
	data io.Reader
}

func NewBlobObject(repo *git.Repository, data io.Reader) *blobObject {
	return &blobObject{
		repo: repo,
		data: data,
	}
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
