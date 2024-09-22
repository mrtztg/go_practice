package findallduplicatenumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllDuplicateNumbers(t *testing.T) {
	testCases := []struct {
		input, output []int
	}{
		{[]int{3, 4, 4, 5, 5}, []int{4, 5}},
		{[]int{5, 4, 7, 2, 3, 5, 3}, []int{3, 5}},
	}
	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.output, findallduplicatenumbers(tc.input))
	}
}
