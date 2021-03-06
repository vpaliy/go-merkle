package merklelib

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
)

func toBytes(item interface{}) []byte {
	switch value := item.(type) {
	case string:
		return []byte(value)
	case []byte:
		return value
	case MerkleNode:
		return value.hashVal
	}
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(item)
	// TODO: check this error here
	if err != nil {
		return nil
	}
	return buf.Bytes()
}

func toHex(src []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func reverse(a ...interface{}) []interface{} {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}
