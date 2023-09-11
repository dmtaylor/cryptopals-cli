package distance

import (
	"errors"
	"math/bits"
)

var MismatchedInputs = errors.New("mismatched inputs for distance")

func BitwiseHamming(in1, in2 []byte) (int, error) {
	if len(in1) != len(in2) {
		return 0, MismatchedInputs
	}
	diff := 0
	for i := 0; i < len(in1); i++ {
		diff += 8 - bits.OnesCount8(^(in1[i] ^ in2[i]))
	}

	return diff, nil
}
