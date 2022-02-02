package object

import (
	"io"
	"io/ioutil"

	"github.com/chrillux/go-wyag/git"
)

type commitObject struct {
	repo *git.Repository
	data io.Reader
	kvlm *KVLM
}

func NewCommitObject(repo *git.Repository, data io.Reader) *commitObject {
	return &commitObject{
		repo: repo,
		data: data,
	}
}

func (o *commitObject) Deserialize(data io.Reader) {
	o.kvlm = ParseKeyValueListWithMessage(data)
}

func (o *commitObject) Serialize() io.Reader {
	return KeyValueListWithMessageSerialize(*o.kvlm)
}

func (o *commitObject) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *commitObject) GetParents() []string {
	o.Deserialize(o.data)
	p := []string{}
	for _, kv := range o.kvlm.KeyValues {
		if kv.Key == "parent" {
			p = append(p, kv.Value)
		}
	}
	return p
}
