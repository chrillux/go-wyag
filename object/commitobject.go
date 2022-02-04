package object

import (
	"io"
	"io/ioutil"
)

type Commit struct {
	data io.Reader
	kvlm *KVLM
}

func NewCommitObject(data io.Reader) *Commit {
	return &Commit{
		data: data,
	}
}

func (o *Commit) Deserialize(data io.Reader) {
	o.kvlm = ParseKeyValueListWithMessage(data)
}

func (o *Commit) Serialize() io.Reader {
	return KeyValueListWithMessageSerialize(*o.kvlm)
}

func (o *Commit) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *Commit) GetObjType() string {
	return "commit"
}

func (o *Commit) GetParents() []string {
	o.Deserialize(o.data)
	p := []string{}
	for _, kv := range o.kvlm.KeyValues {
		if kv.Key == "parent" {
			p = append(p, kv.Value)
		}
	}
	return p
}

func (o *Commit) KVLM() *KVLM {
	o.Deserialize(o.data)
	return o.kvlm
}
