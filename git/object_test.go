package git

import (
	"fmt"
	"strings"
	"testing"
)

// func TestObject_deserialize(t *testing.T) {
// 	type fields struct {
// 		serializedData   []byte
// 		deserializedData []byte
// 		dataLen          int
// 		hash             string
// 		objType          string
// 	}
// 	type args struct {
// 		data io.Reader
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "deserialize",
// 			// fields: fields{
// 			// 	serializedData: []byte{98, 108, 111, 98, 32, 55, 0, 102, 111, 111, 98, 97, 114, 10},
// 			// 	hash:           "323fae03f4606ea9991df8befbb2fca795e648fa",
// 			// 	objType:        "blob",
// 			// 	dataLen:        7,
// 			// },
// 			args: args{
// 				data: bytes.NewReader([]byte{98, 108, 111, 98, 32, 55, 0, 102, 111, 111, 98, 97, 114, 10}),
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			o := &Object{}
// 			if err := o.Deserialize(tt.args.data); (err != nil) != tt.wantErr {
// 				t.Errorf("Object.deserialize() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			// assert.Equal(t, fmt.Sprintf("%s\n", "foobar"), )
// 		})
// 	}
// }

func TestParseKeyValueListWithMessage(t *testing.T) {
	data := `tree 3f1ef6b311330cd0571e1d4755d0fc35bb4d66c3
committer GitHub <noreply@github.com> 1634567052 +0200
 tjolaho svejsan hej
gpgsig -----BEGIN PGP SIGNATURE-----
 
 wsBcBAABCAAQBQJhbYOMCRBK7hj4Ov3rIwAAEX8IAFo1+qOvEwYFm/WYf/LfsCPp
 BzbkI2LBEtfQbNebIy3aI7CSQZOd4A+RSHPGxIBjWcAlJ9hcZgKobyWhqGIbmTuL
 18KqXsuOe+0u/ObgzRow51WvjRrRjP20ORJ6DxTsao+IiDCQP1MH1jojJTqhfTjX
 7y2RCW53XG6OwfFtq60Ftzx6uaUTC6ed2yMwTTly+lyV46h8vTRIq0joKbQgF6Ap
 N6Myy1h/VcOtX0B8cYKtg9DsDTdluTdsqhSKI5lp/2wfncnHU599VlROWten6jKR
 5WO/rEEOjxC8f9GQCloLyj1i3bXWtBzuX/hxP9M95OcsGTZvnKAWnd/8pEwhpmc=
 =jJhI
 -----END PGP SIGNATURE-----
 

QRND-7659 (#11377)

* QRND-7659

* QRND-7659`
	r := strings.NewReader(data)
	kvlm := ParseKeyValueListWithMessage(r)
	fmt.Printf("%v", kvlm)
}
