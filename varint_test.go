package sqlite3utils

import (
	"testing"
)

func assertEq(t *testing.T, expect, actual byte) {
	if expect != actual {
		t.Error("Expected:", expect, "Actual:", actual)
	}
}

func assertEqAll(t *testing.T, expect byte, actual []byte) {

	for i, a := range actual {
		if expect != a {
			t.Error("Expected", expect, "Actual", a, "at", i, "in", actual)
		}
	}

}

func TestEncodeMax(t *testing.T) {

	t.Run("v=240", func(t *testing.T) {
		bytes := Encode(240)
		//fmt.Println(bytes)
		assertEq(t, 240, bytes[0])
	})

	testCase := map[uint64]byte{
		2287:                 248,
		67823:                249,
		16777215:             250,
		4294967295:           251,
		1099511627775:        252,
		281474976710655:      253,
		72057594037927935:    254,
		18446744073709551615: 255,
	}

	for k, v := range testCase {
		t.Run("v="+string(k), func(t *testing.T) {
			bytes := Encode(k)
			//fmt.Println(v, bytes)
			assertEq(t, v, bytes[0])
			assertEqAll(t, 255, bytes[1:])
		})
	}
}

func TestDencodeMax(t *testing.T) {

	testCase := map[uint64][]byte{}
	testCase[240] = []byte{240}
	testCase[2287] = []byte{248, 255}
	testCase[67823] = []byte{249, 255, 255}
	testCase[16777215] = []byte{250, 255, 255, 255}
	testCase[4294967295] = []byte{251, 255, 255, 255, 255}
	testCase[1099511627775] = []byte{252, 255, 255, 255, 255, 255}
	testCase[281474976710655] = []byte{253, 255, 255, 255, 255, 255, 255}
	testCase[72057594037927935] = []byte{254, 255, 255, 255, 255, 255, 255, 255}
	testCase[18446744073709551615] = []byte{255, 255, 255, 255, 255, 255, 255, 255, 255}

	for k, v := range testCase {
		//fmt.Println(Decode([]byte{240}))

		n, _ := Decode(v)
		if n != k {
			t.Error("n>", k, n)
		}
	}
}
