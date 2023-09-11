package distance

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitwiseHamming(t *testing.T) {
	type args struct {
		in1 []byte
		in2 []byte
	}
	tests := []struct {
		name        string
		args        args
		want        int
		expectedErr error
	}{
		{
			"empty",
			args{
				[]byte{},
				[]byte{},
			},
			0,
			nil,
		},
		{
			"mismatched_lengths",
			args{
				[]byte{0x1, 0x2},
				[]byte{0x2},
			},
			0,
			MismatchedInputs,
		},
		{
			"cryptopals_example",
			args{
				[]byte("this is a test"),
				[]byte("wokka wokka!!!"),
			},
			37,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitwiseHamming(tt.args.in1, tt.args.in2)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error(), "mismatching errors")
			} else {
				if assert.NoError(t, err, "got unexpected error") {
					assert.Equal(t, tt.want, got, "mismatched results")
				}
			}
		})
	}
}
