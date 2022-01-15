package git

import (
	"io"
)

func CatFile(f io.Reader) (*Object, error) {
	o := &Object{}
	err := o.deserialize(f)
	if err != nil {
		return nil, err
	}
	return o, nil
}
