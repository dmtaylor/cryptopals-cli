package xor

import "github.com/dmtaylor/cryptopals-cli/ciphers"

func FixedXor(in, key []byte) ([]byte, error) {
	if len(in) != len(key) {
		return nil, ciphers.InvalidKeyLength
	}
	result := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		result[i] = in[i] ^ key[i]
	}
	return result, nil
}
