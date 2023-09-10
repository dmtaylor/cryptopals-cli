package encoding

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type InvalidEncodingError struct {
	badValue rune
}

func (e InvalidEncodingError) Error() string {
	return fmt.Sprintf("invalid character in encoding %c", e.badValue)
}

var InvalidHexLengthError = errors.New("invalid length of hex encoding")

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
	if len(s)%2 == 1 {
		return nil, InvalidHexLengthError
	}
	result := make([]byte, 0, len(s)/2)
	runes := []rune(s)
	for i := 0; i < len(runes); i += 2 {
		ho, ok := hexRuneToBin[runes[i]]
		if !ok {
			return nil, InvalidEncodingError{runes[i]}
		}
		lo, ok := hexRuneToBin[runes[i+1]]
		if !ok {
			return nil, InvalidEncodingError{runes[i+1]}
		}
		result = append(result, (ho<<4)|lo)
	}

	return result, nil
}

func BinToHex(r io.Reader) (string, error) {
	output := strings.Builder{}
	buf := make([]byte, 64)
	for {
		n, err := r.Read(buf)
		for i := 0; i < n; i++ {
			ho, ok := hexBinToRune[buf[i]>>4]
			if !ok {
				return "", fmt.Errorf("bad map value %b", buf[i]>>4)
			}
			output.WriteRune(ho)
			lo, ok := hexBinToRune[buf[i]&0x0F]
			if !ok {
				return "", fmt.Errorf("bad map value %b", buf[i]&0x0F)
			}
			output.WriteRune(lo)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return output.String(), fmt.Errorf("read error: %w", err)
		}
	}
	return output.String(), nil
}
