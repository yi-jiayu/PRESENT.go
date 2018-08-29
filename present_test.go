package present_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yi-jiayu/PRESENT.go"
)

var cases = []struct {
	Key        string
	Plaintext  string
	Ciphertext string
}{
	{
		Key:        "00000000000000000000",
		Plaintext:  "0000000000000000",
		Ciphertext: "5579C1387B228445",
	}, {
		Key:        "FFFFFFFFFFFFFFFFFFFF",
		Plaintext:  "0000000000000000",
		Ciphertext: "E72C46C0F5945049",
	},
	{
		Key:        "00000000000000000000",
		Plaintext:  "FFFFFFFFFFFFFFFF",
		Ciphertext: "A112FFC72F68417B",
	}, {
		Key:        "FFFFFFFFFFFFFFFFFFFF",
		Plaintext:  "FFFFFFFFFFFFFFFF",
		Ciphertext: "3333DCD3213210D2",
	},
}

func decodeHex(s string) []byte {
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		panic(err)
	}
	return dst
}

func TestNewCipher(t *testing.T) {
	t.Parallel()

	t.Run("invalid key size", func(t *testing.T) {
		var key []byte
		_, err := present.NewCipher(key)
		if err == nil {
			t.Fail()
		}
		assert.Equal(t, present.KeySizeError(0), err)
		assert.Equal(t, "present: invalid key size 0", err.Error())
	})
}

func TestBlock_Encrypt(t *testing.T) {
	t.Parallel()

	for _, c := range cases {
		key := decodeHex(c.Key)
		cipher, err := present.NewCipher(key)
		if err != nil {
			t.Fatal(err)
		}

		plaintext := decodeHex(c.Plaintext)
		dst := make([]byte, cipher.BlockSize())
		cipher.Encrypt(dst, plaintext)

		ciphertext := decodeHex(c.Ciphertext)
		assert.Equal(t, ciphertext, dst)
	}
}

func TestBlock_Decrypt(t *testing.T) {
	t.Parallel()

	for _, c := range cases {
		key := decodeHex(c.Key)
		cipher, err := present.NewCipher(key)
		if err != nil {
			t.Fatal(err)
		}

		ciphertext := decodeHex(c.Ciphertext)
		dst := make([]byte, cipher.BlockSize())
		cipher.Decrypt(dst, ciphertext)

		plaintext := decodeHex(c.Plaintext)
		assert.Equal(t, plaintext, dst)
	}
}
