package sys

import (
	"hash"
)

type Bloomfilter struct {
	Bytes  []byte
	Hashes []hash.Hash64
}

func (b Bloomfilter) Push(key string) {
	var byteLen = len(b.Bytes)
	for _, v := range b.Hashes {
		v.Reset()
		v.Write([]byte(key))
		res := v.Sum64()
		var yByte = res % uint64(byteLen)
		var yBit = res % 7
		b.Bytes[yByte] |= 1 << yBit
	}
}

func (b Bloomfilter) Exists(key string) bool {
	var byteLen = len(b.Bytes)
	for _, v := range b.Hashes {
		v.Reset()
		v.Write([]byte(key))
		res := v.Sum64()
		var yByte = res % uint64(byteLen)
		var yBit = res % 7
		if b.Bytes[yByte]|1<<yBit != b.Bytes[yByte] {
			return false
		}
	}
	return true
}
