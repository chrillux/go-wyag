package object

import (
	"fmt"
	"strings"
	"testing"
)

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
