package encoding

import (
	"fmt"
	"io"
)

type InvalidEncodingError struct {
	badValue rune
}

func (e InvalidEncodingError) Error() string {
	return fmt.Sprintf("invalid character in encoding %c", e.badValue)
}

// BinToBase64 encodes binary stream in base64
func BinToBase64(r io.Reader) (string, error) {
	// TODO implement

	// TODO remove stub return
	return "", nil
}

// Base64ToBin reads base64 string and converts to binary
func Base64ToBin(s string) ([]byte, error) {
	// TODO implement

	// TODO remove stub return
	return []byte{}, nil
}

func HexToBin(s string) ([]byte, error) {
	// TODO implement

	// TODO remove stub return
	return []byte{}, nil
}

func BinToHex(r io.Reader) (string, error) {
	// TODO implement

	// TODO remove stub return
	return "", nil
}
