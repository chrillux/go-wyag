package object

import (
	"io"
	"log"
	"reflect"
	"testing"
)

func getData() io.Reader {
	treehash := "02a32a7bc48d64a5b9f04aa5aeea5d91929865d0"
	o, err := ReadObject(treehash)
	if err != nil {
		log.Fatalf("error reading object: %v", err)
	}
	return o.obj.Serialize()
}

func Test_treeParseOne(t *testing.T) {
	data := getData()
	type args struct {
		raw io.Reader
		pos int
	}
	tests := []struct {
		name  string
		args  args
		want  gitTreeLeaf
		want1 int
	}{
		{
			name: "Read one leaf",
			args: args{
				raw: data,
				pos: 0,
			},
			want: gitTreeLeaf{
				mode: "100644",
				path: "catfile.go",
				sha:  "ef51f132cfe02e71cda4638f01c190cc91e035cc",
			},
			want1: 38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := treeParseOne(tt.args.raw, tt.args.pos)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeParseOne() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("treeParseOne() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_treeParse(t *testing.T) {
	type args struct {
		raw io.Reader
	}
	tests := []struct {
		name string
		args args
		want []gitTreeLeaf
	}{
		{
			name: "Read whole tree object",
			args: args{
				raw: getData(),
			},
			want: []gitTreeLeaf{
				{
					mode: "100644",
					path: "catfile.go",
					sha:  "ef51f132cfe02e71cda4638f01c190cc91e035cc",
				},
				{
					mode: "100644",
					path: "hashobject.go",
					sha:  "87bb585a4dc1438365ff495dbb67493be9662804",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeParse(tt.args.raw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_treeSerialize(t *testing.T) {
	type args struct {
		gtls []gitTreeLeaf
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Serialize a tree object",
			args: args{
				gtls: []gitTreeLeaf{
					{
						mode: "100644",
						path: "foo.sh",
						sha:  "02a32a7bc48d64a5b9f04aa5aeea5d91929865d0",
					},
				},
			},
			want: []byte{
				49, 48, 48, 54, 52, 52, 32, 102, 111, 111, 46, 115, 104, 0, 2, 163, 42, 123, 196, 141, 100, 165, 185, 240, 74, 165, 174, 234, 93, 145, 146, 152, 101, 208,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := treeSerialize(tt.args.gtls); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("treeSerialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
