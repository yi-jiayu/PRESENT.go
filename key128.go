package present

type key128 struct {
	A, B uint64
}

func (k *key128) copy() key {
	cpy := *k
	return &cpy
}

func (k *key128) rotate() {
	panic("implement me")
}

func (k *key128) sBox() {
	panic("implement me")
}

func (k *key128) xor(ctr uint64) {
	panic("implement me")
}

func (k *key128) roundKey() uint64 {
	panic("implement me")
}

func newKey128(key []byte) *key128 {
	A, B := decompose(key)
	return &key128{A, B}
}
