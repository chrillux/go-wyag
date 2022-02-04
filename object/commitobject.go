package object

import (
	"io"
	"io/ioutil"

	"github.com/chrillux/go-wyag/git"
)

type CommitObject struct {
	repo *git.Repository
	data io.Reader
	kvlm *KVLM
}

func NewCommitObject(repo *git.Repository, data io.Reader) *CommitObject {
	return &CommitObject{
		repo: repo,
		data: data,
	}
}

func (o *CommitObject) Deserialize(data io.Reader) {
	o.kvlm = ParseKeyValueListWithMessage(data)
}

func (o *CommitObject) Serialize() io.Reader {
	return KeyValueListWithMessageSerialize(*o.kvlm)
}

func (o *CommitObject) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *CommitObject) GetObjType() string {
	return "commit"
}

func (o *CommitObject) GetParents() []string {
	o.Deserialize(o.data)
	p := []string{}
	for _, kv := range o.kvlm.KeyValues {
		if kv.Key == "parent" {
			p = append(p, kv.Value)
		}
	}
	return p
}

func (o *CommitObject) KVLM() *KVLM {
	o.Deserialize(o.data)
	return o.kvlm
}
