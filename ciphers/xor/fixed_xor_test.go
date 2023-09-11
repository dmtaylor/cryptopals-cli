package xor

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/dmtaylor/cryptopals-cli/ciphers"
	"github.com/dmtaylor/cryptopals-cli/util/encoding"
)

func TestFixedXor(t *testing.T) {
	type args struct {
		in  []byte
		key []byte
	}
	tests := []struct {
		name          string
		args          args
		want          []byte
		expectedError error
	}{
		{
			"empty",
			args{
				[]byte{},
				[]byte{},
			},
			[]byte{},
			nil,
		},
		{
			"basic_case",
			args{
				[]byte{0x03, 0x03},
				[]byte{0x01, 0x02},
			},
			[]byte{0x02, 0x01},
			nil,
		},
		{
			"invalid_keylen",
			args{
				[]byte{0, 1, 2},
				[]byte{1},
			},
			nil,
			ciphers.InvalidKeyLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FixedXor(tt.args.in, tt.args.key)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "mismatching errors")
			} else {
				assert.Equal(t, tt.want, got, "mismatching binary encoding")
			}
		})
	}
}

func TestFixedXorCryptopalsValidation(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	key := "686974207468652062756c6c277320657965"
	inBin, err := encoding.HexToBin(input)
	require.Nil(t, err, "decoding got error")
	keyBin, err := encoding.HexToBin(key)
	require.Nil(t, err, "decoding key got error")
	got, err := FixedXor(inBin, keyBin)
	if assert.NoError(t, err, "xor failed") {
		res, err := encoding.BinToHex(bytes.NewReader(got))
		require.Nil(t, err, "encoding result failed")
		assert.Equal(t, "746865206b696420646f6e277420706c6179", res, "mismatching results")
	}

}
