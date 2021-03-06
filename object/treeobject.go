package object

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

type Tree struct {
	data  io.Reader
	items []gitTreeLeaf
}

// gitTreeLeaf is a single tree record, i.e a single path or file.
type gitTreeLeaf struct {
	mode string
	path string
	sha  string
}

func (g *gitTreeLeaf) Mode() string {
	return g.mode
}
func (g *gitTreeLeaf) Path() string {
	return g.path
}
func (g *gitTreeLeaf) SHA() string {
	return g.sha
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

func treeSerialize(gtls []gitTreeLeaf) io.Reader {
	ret := []byte{}
	for _, gtl := range gtls {
		ret = append(ret, []byte(gtl.mode)...)
		ret = append(ret, byte(32))
		ret = append(ret, []byte(gtl.path)...)
		ret = append(ret, byte(0))
		shabytes, _ := hex.DecodeString(gtl.sha)
		ret = append(ret, shabytes...)
	}
	return bytes.NewReader(ret)
}

func NewTreeObject(data io.Reader) *Tree {
	to := &Tree{
		data: data,
	}
	to.Deserialize(to.data)
	return to
}

func (o *Tree) Deserialize(data io.Reader) {
	o.items = treeParse(data)
}

func (o *Tree) Serialize() io.Reader {
	return treeSerialize(o.items)
}

func (o *Tree) String() string {
	var treeobjAsString string
	for _, item := range o.Items() {
		treeobjAsString += fmt.Sprintf("%s\n", item)
	}
	return treeobjAsString
}

func (o *Tree) GetObjType() string {
	return "tree"
}

func (o *Tree) Items() []gitTreeLeaf {
	return o.items
}
