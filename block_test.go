package present

import (
	"fmt"
	"os"
	"testing"
)

func Test_sBoxLayer(t *testing.T) {
	t.Run("s box layer", func(t *testing.T) {
		var state uint64 = 0x123456789abcdef
		var expected uint64 = 0xc56b90ad3ef84712
		actual := sBoxLayer(state, sBox)
		if actual != expected {
			t.Fail()
		}
	})
	t.Run("inverse s box layer", func(t *testing.T) {
		var state uint64 = 0xc56b90ad3ef84712
		var expected uint64 = 0x123456789abcdef
		actual := sBoxLayer(state, sBoxInv)
		if actual != expected {
			t.Fail()
		}
	})
}

func Test_pLayer(t *testing.T) {
	t.Run("p layer", func(t *testing.T) {
		var state uint64 = 0xaaaaaaaaaaaaaaaa
		var expected uint64 = 0xffff0000ffff0000
		actual := pLayer(state, p)
		if actual != expected {
			fmt.Fprintf(os.Stderr, "expected: %x\ngot     : %x\n", expected, actual)
			t.Fail()
		}
	})
	t.Run("inverse p layer", func(t *testing.T) {
		var state uint64 = 0xffff0000ffff0000
		var expected uint64 = 0xaaaaaaaaaaaaaaaa
		actual := pLayer(state, pInv)
		if actual != expected {
			fmt.Fprintf(os.Stderr, "expected: %x\ngot     : %x\n", expected, actual)
			t.Fail()
		}
	})
}
