package object

import (
	"io"
	"log"
	"reflect"
	"testing"
)

func Test_treeParseOne(t *testing.T) {
	treehash := "02a32a7bc48d64a5b9f04aa5aeea5d91929865d0"
	o, err := ReadObject(treehash)
	if err != nil {
		log.Fatalf("error reading object: %v", err)
	}
	data := o.obj.Serialize()
	type args struct {
		raw io.Reader
		pos int
	}
	tests := []struct {
		name  string
		args  args
		want  *gitTreeLeaf
		want1 int
	}{
		{
			name: "Read one leaf",
			args: args{
				raw: data,
				pos: 0,
			},
			want: &gitTreeLeaf{
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
