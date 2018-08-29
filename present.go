package present

import (
	"crypto/cipher"
	"strconv"
)

type KeySizeError int

func (k KeySizeError) Error() string {
	return "present: invalid key size " + strconv.Itoa(int(k))
}

// NewCipher creates a new cipher.Block.
// The argument should be the PRESENT key,
// which is either 10 or 16 bytes long
// for key lengths of 80 bits and 128 bits respectively.
func NewCipher(key []byte) (b cipher.Block, err error) {
	k := len(key)
	switch k {
	default:
		err = KeySizeError(k)
	case 10:
		b = &block{
			Key: newKey80(key),
		}
	case 16:
		b = &block{
			Key: newKey128(key),
		}
	}
	return
}
