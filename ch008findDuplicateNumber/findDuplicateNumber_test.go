package ch008findduplicatenumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{[]int{1, 4, 4, 3, 2}, 4},
		{[]int{2,1,3,3,5,4}, 3},
		{[]int{2,4,1,4,4}, 4},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.output, findDuplicateNumber(tc.input))
	}
}
