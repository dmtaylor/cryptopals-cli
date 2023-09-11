package xor

func SingleByteXor(in []byte, key byte) []byte {
	result := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		result[i] = in[i] ^ key
	}
	return result
}
