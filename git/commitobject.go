package git

import (
	"io"
	"io/ioutil"
)

type commitObject struct {
	repo    *Repository
	data    io.Reader
	kvlm    *KVLM
	objType string
}

func NewCommitObject(repo *Repository, data io.Reader) *commitObject {
	o := &commitObject{
		repo:    repo,
		data:    data,
		kvlm:    ParseKeyValueListWithMessage(data),
		objType: "commit",
	}
	return o
}

func (o *commitObject) Deserialize(data io.Reader) {
	o.data = data
}

func (o *commitObject) Serialize() io.Reader {
	return o.data
}

func (o *commitObject) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *commitObject) GetParents() []string {
	p := []string{}
	for _, kv := range o.kvlm.KeyValues {
		if kv.Key == "parent" {
			p = append(p, kv.Value)
		}
	}
	return p
}
