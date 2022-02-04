package object

import (
	"io"
	"io/ioutil"
)

type blobObject struct {
	data io.Reader
}

func NewBlobObject(data io.Reader) *blobObject {
	return &blobObject{
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

func (o *blobObject) GetObjType() string {
	return "blob"
}
