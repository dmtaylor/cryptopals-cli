package encoding

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64ToBin(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        []byte
		expectedErr error
	}{
		{
			"empty",
			"",
			[]byte{},
			nil,
		},
		{
			"basic_wiki_example",
			"TWFu",
			[]byte{0x4d, 0x61, 0x6e},
			nil,
		},
		{
			"wiki_padding_1",
			"TWE=",
			[]byte{0x4d, 0x61},
			nil,
		},
		{
			"wiki_padding_2",
			"TQ==",
			[]byte{0x4d},
			nil,
		},
		{
			"invalid_length",
			"ads",
			nil,
			InvalidLengthError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Base64ToBin(tt.input)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error(), "mismatching errors")
			} else {
				assert.Equal(t, tt.want, got, "mismatching binary encoding")
			}
		})
	}
}

func TestBinToBase64(t *testing.T) {
	tests := []struct {
		name          string
		input         io.Reader
		want          string
		expectedError error
	}{
		{
			"empty",
			bytes.NewReader([]byte{}),
			"",
			nil,
		},
		{
			"wiki_example",
			bytes.NewReader([]byte("Many hands make light work.")),
			"TWFueSBoYW5kcyBtYWtlIGxpZ2h0IHdvcmsu",
			nil,
		},
		{
			"padded_example_1",
			bytes.NewReader([]byte("Ma")),
			"TWE=",
			nil,
		},
		{
			"padded_example_2",
			bytes.NewReader([]byte("M")),
			"TQ==",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinToBase64(tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "mismatching errors")
			} else {
				assert.Equal(t, tt.want, got, "mismatching encoding")
			}
		})
	}
}

func TestBinToHex(t *testing.T) {
	bigBytes := make([]byte, 65)
	bigEncoding := strings.Builder{}
	for i := 0; i < len(bigBytes); i++ {
		bigBytes[i] = 1
		bigEncoding.WriteString("01")
	}
	tests := []struct {
		name          string
		input         io.Reader
		want          string
		expectedError error
	}{
		{
			"empty",
			bytes.NewReader([]byte{}),
			"",
			nil,
		},
		{
			"simple",
			bytes.NewReader([]byte{0x4, 0xAF, 0x64}),
			"04af64",
			nil,
		},
		{
			"multiple_loops",
			bytes.NewReader(bigBytes),
			bigEncoding.String(),
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinToHex(tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "mismatching errors")
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.want, got, "mismatching decoding")
				}
			}
		})
	}
}

func TestHexToBin(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		want          []byte
		expectedError error
	}{
		{
			"empty",
			"",
			[]byte{},
			nil,
		},
		{
			"basic",
			"01af32",
			[]byte{0x1, 0xAF, 0x32},
			nil,
		},
		{
			"invalid_input_length",
			"012",
			[]byte{},
			InvalidLengthError,
		},
		{
			"invalid_encoding1",
			"01g4",
			nil,
			InvalidEncodingError{'g'},
		},
		{
			"invalid_encoding2",
			"014g",
			nil,
			InvalidEncodingError{'g'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToBin(tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "mismatching errors")
			} else {
				if assert.NoError(t, err) {
					assert.Equal(t, tt.want, got, "mismatching decoding")
				}
			}
		})
	}
}

func TestInvalidEncodingError_Error(t *testing.T) {
	tests := []struct {
		name     string
		badValue rune
		want     string
	}{
		{
			"basic",
			'\\',
			"invalid character in encoding \\",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := InvalidEncodingError{
				badValue: tt.badValue,
			}
			assert.EqualError(t, e, tt.want)
		})
	}
}

func TestCryptopalsValidation(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	binary, err := HexToBin(input)
	if assert.NoError(t, err, "decoding failed") {
		got, err := BinToBase64(bytes.NewReader(binary))
		if assert.NoError(t, err, "encoding failed") {
			assert.Equal(t, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", got, "decoded value doesn't match")
		}
	}
}

func TestCryptopalsValidationInverse(t *testing.T) {
	input := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	binary, err := Base64ToBin(input)
	if assert.NoError(t, err, "decoding failed") {
		got, err := BinToHex(bytes.NewReader(binary))
		if assert.NoError(t, err, "encoding failed") {
			assert.Equal(t, "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", got, "decoded value doesn't match")
		}
	}
}
