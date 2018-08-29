package present

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_key80_rotate(t *testing.T) {
	t.Parallel()

	k := &key80{
		A: 0xc,
	}
	k.rotate()
	expected := key80{
		A: 1 << 63,
		B: 1 << 48,
	}
	assert.Equal(t, expected, *k)
}

func Test_key80_update(t *testing.T) {
	t.Parallel()

	k := &key80{}
	updateKey(k, 1)
	expected := key80{
		A: 0x3 << 62,
		B: 1 << 63,
	}
	assert.Equal(t, expected, *k)
}
