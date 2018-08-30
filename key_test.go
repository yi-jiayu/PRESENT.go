package present

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decompose(t *testing.T) {
	cases := []struct {
		key []byte
		A   uint64
		B   uint64
	}{
		{
			key: []byte{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
			A:   1 << 16,
		},
		{
			key: []byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			A:   1,
		},
		{
			key: []byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
			A:   1,
			B:   1 << 48,
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			A, B := decompose(c.key)
			assert.Equal(t, c.A, A)
			assert.Equal(t, c.B, B)
		})
	}
}
