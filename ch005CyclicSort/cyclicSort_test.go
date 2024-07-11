package ch005CyclicSort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	s := Solution{}

	testCases := []struct {
		input, output []int
	}{
		{
			input:  []int{3, 1, 5, 4, 2},
			output: []int{1, 2, 3, 4, 5},
		}, {
			input:  []int{2, 6, 4, 3, 1, 5},
			output: []int{1, 2, 3, 4, 5, 6},
		}, {
			input:  []int{1, 5, 6, 4, 3, 2},
			output: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.output, s.sort(tc.input), "input:%+v", tc.input)
	}
}
