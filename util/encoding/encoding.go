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

var InvalidLengthError = errors.New("invalid length of encoding")

// BinToBase64 encodes binary stream in base64
func BinToBase64(r io.Reader) (string, error) {
	result := strings.Builder{}
	buf := make([]byte, 3)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			result.WriteRune(b64BinToRune[buf[0]>>2])
			if n == 1 {
				result.WriteRune(b64BinToRune[(buf[0]&03)<<4])
				result.WriteRune(padding)
				result.WriteRune(padding)
			} else if n == 2 {
				t := ((buf[0] & 03) << 4) | buf[1]>>4
				result.WriteRune(b64BinToRune[t])
				r := (buf[1] & 0b1111) << 2
				result.WriteRune(b64BinToRune[r])
				result.WriteRune(padding)
			} else if n == 3 {
				t := ((buf[0] & 03) << 4) | buf[1]>>4
				result.WriteRune(b64BinToRune[t])
				r := ((buf[1] & 0b1111) << 2) | ((buf[2] & 0b11000000) >> 6)
				result.WriteRune(b64BinToRune[r])
				l := buf[2] & 0b111111
				result.WriteRune(b64BinToRune[l])
			}
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return result.String(), err
		}
	}
	return result.String(), nil
}

// Base64ToBin reads base64 string and converts to binary
func Base64ToBin(s string) ([]byte, error) {
	if len(s)%4 != 0 {
		return nil, InvalidLengthError
	}
	result := make([]byte, 0, len(s)*3/4)
	runes := []rune(s)
	for i := 0; i < len(runes); i += 4 {
		f := (b64RuneToBin[runes[i]] << 2) | ((b64RuneToBin[runes[i+1]] & 0b110000) >> 4)
		result = append(result, f)
		if runes[2] == padding {
			break
		}
		t := ((b64RuneToBin[runes[i+1]] & 0b1111) << 4) | ((b64RuneToBin[runes[i+2]] & 0xFC) >> 2)
		result = append(result, t)
		if runes[3] == padding {
			break
		}
		l := ((b64RuneToBin[runes[i+2]] & 0b11) << 6) | b64RuneToBin[runes[i+3]]
		result = append(result, l)
	}
	return result, nil
}

// HexToBin converts hex stream into represented byte slice
func HexToBin(s string) ([]byte, error) {
	if len(s)%2 == 1 {
		return nil, InvalidLengthError
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

// BinToHex encodes data from reader into hex representation
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
