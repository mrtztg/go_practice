package ch007findAllMissingNumbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllMissingNumbers(t *testing.T) {
	s := solution{}
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			[]int{2, 3, 1, 8, 2, 3, 5, 1},
			[]int{4, 6, 7},
		}, {
			[]int{2, 4, 1, 2},
			[]int{3},
		}, {
			[]int{2, 3, 2, 1},
			[]int{4},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.output, s.findAllMissingNumbers(tc.input))
	}
}
