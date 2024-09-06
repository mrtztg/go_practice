package ch006findMissingNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMissingNumber(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{4, 0, 3, 1},
			output: 2,
		}, {
			input:  []int{8, 3, 5, 2, 4, 6, 0, 1},
			output: 7,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.output, s.findMissingNumber(tc.input))
	}
}
