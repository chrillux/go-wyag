package object

import (
	"bytes"
	"encoding/hex"
	"io"
	"io/ioutil"

	"github.com/chrillux/go-wyag/git"
)

type treeObject struct {
	repo *git.Repository
	data io.Reader
}

// gitTreeLeaf is a single tree record, i.e a single path or file.
type gitTreeLeaf struct {
	mode string
	path string
	sha  string
}

// treeParseOne a tree is a concatenation of records of the format: [mode] space [path] 0x00 [sha-1]
func treeParseOne(raw io.Reader, pos int) (*gitTreeLeaf, int) {
	data, _ := ioutil.ReadAll(raw)
	space := bytes.IndexByte(data[pos:], 32)
	mode := string(data[pos:space])
	null := bytes.IndexByte(data[pos:], 0)
	path := string(data[space+1 : null])
	sha := hex.EncodeToString(data[null+1 : null+21])

	return &gitTreeLeaf{
		mode: mode,
		path: path,
		sha:  sha,
	}, pos + null + 21
}

func NewTreeObject(repo *git.Repository, data io.Reader) *treeObject {
	o := &treeObject{
		repo: repo,
		data: data,
	}
	return o
}

func (o *treeObject) Deserialize(data io.Reader) {
	// tbd
}

func (o *treeObject) Serialize() io.Reader {
	return o.data
}

func (o *treeObject) String() string {
	s, _ := ioutil.ReadAll(o.Serialize())
	return string(s)
}

func (o *treeObject) GetParents() []string {
	return nil
}
