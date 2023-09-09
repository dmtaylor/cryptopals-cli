package encoding

import (
	"bytes"
	"io"
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
		// TODO: Add test cases.
		{
			"empty",
			"",
			[]byte{},
			nil,
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
		// TODO: Add test cases.
		{
			"empty",
			bytes.NewReader([]byte{}),
			"",
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
	tests := []struct {
		name          string
		input         io.Reader
		want          string
		expectedError error
	}{
		// TODO: Add test cases.
		{
			"empty",
			bytes.NewReader([]byte{}),
			"",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinToHex(tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "mismatching errors")
			} else {
				assert.Equal(t, tt.want, got, "mismatching decoding")
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
		// TODO: Add test cases.
		{
			"empty",
			"",
			[]byte{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToBin(tt.input)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "mismatching errors")
			} else {
				assert.Equal(t, tt.want, got, "mismatching decoding")
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
