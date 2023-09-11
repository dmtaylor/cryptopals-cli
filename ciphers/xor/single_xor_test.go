package xor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleByteXor(t *testing.T) {
	type args struct {
		in  []byte
		key byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"empty",
			args{
				[]byte{},
				0x0,
			},
			[]byte{},
		},
		{
			"basic",
			args{
				[]byte{0x00, 0x01, 0x02},
				0x01,
			},
			[]byte{0x01, 0x00, 0x03},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SingleByteXor(tt.args.in, tt.args.key))
		})
	}
}
