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
func treeParseOne(raw io.Reader, pos int) (gitTreeLeaf, int) {
	data, _ := ioutil.ReadAll(raw)
	space := bytes.IndexByte(data[pos:], 32) + pos
	mode := string(data[pos:space])
	null := bytes.IndexByte(data[pos:], 0) + pos
	path := string(data[space+1 : null])
	sha := hex.EncodeToString(data[null+1 : null+21])

	return gitTreeLeaf{
		mode: mode,
		path: path,
		sha:  sha,
	}, pos + null + 21
}

// treeParse parses all lines in a tree commit into a list of tree objects. This is used to deserialize tree data.
func treeParse(raw io.Reader) []gitTreeLeaf {
	data, _ := ioutil.ReadAll(raw)
	pos := 0
	max := len(data)
	treeleafs := []gitTreeLeaf{}
	treeleaf := gitTreeLeaf{}
	for i := pos; i < max; i += pos {
		treeleaf, pos = treeParseOne(bytes.NewReader(data), pos)
		treeleafs = append(treeleafs, treeleaf)
	}
	return treeleafs
}

func treeSerialize(gtls []gitTreeLeaf) []byte {
	ret := []byte{}
	for _, gtl := range gtls {
		ret = append(ret, []byte(gtl.mode)...)
		ret = append(ret, byte(32))
		ret = append(ret, []byte(gtl.path)...)
		ret = append(ret, byte(0))
		shabytes, _ := hex.DecodeString(gtl.sha)
		ret = append(ret, shabytes...)
	}
	return ret
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
