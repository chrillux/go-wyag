package object

import (
	"io"
	"io/ioutil"
)

type Blob struct {
	data io.Reader
}

func NewBlob(data io.Reader) *Blob {
	return &Blob{
		data: data,
	}
}

func (o *Blob) Deserialize(data io.Reader) {
	o.data = data
}

func (o *Blob) Serialize() io.Reader {
	return o.data
}

func (o *Blob) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *Blob) GetObjType() string {
	return "blob"
}
