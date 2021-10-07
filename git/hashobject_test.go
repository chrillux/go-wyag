package git

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestBlobObject_serialize(t *testing.T) {
	rd := []byte("foobar")
	tests := []struct {
		name string
		b    *BlobObject
	}{
		{
			name: "serialize blob",
			b:    &BlobObject{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.b.serialize(rd)
			if err != nil {
				fmt.Printf("error burrr: %v\n", err)
			}
			assert.DeepEqual(t, []byte{98, 108, 111, 98, 32, 54, 0, 102, 111, 111, 98, 97, 114}, tt.b.serializedData)
		})
	}
}

func TestHashObject(t *testing.T) {
	t.Run("hash object", func(t *testing.T) {
		r := New()
		hash, err := r.HashObject("blob", []byte("hejsvejlolboll12345"), false)
		if err != nil {
			fmt.Printf("error burri: %v\n", err)
		}
		assert.Equal(t, "67ca443d6a0dd649f299807492355c044c4c9366", *hash)
	})
}
