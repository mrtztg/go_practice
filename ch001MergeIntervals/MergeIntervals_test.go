package ch001MergeIntervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolutionMerge(t *testing.T) {
	s := solution{}
	testCases := []struct {
		input, output []interval
	}{
		{
			input:  []interval{{1, 4}, {2, 5}, {7, 9}},
			output: []interval{{1, 4}, {7, 9}},
		},
		{
			input:  []interval{{6, 7}, {2, 4}, {5, 9}},
			output: []interval{{2, 4}, {5, 9}},
		},
		{
			input:  []interval{{1, 4}, {2, 6}, {3, 5}},
			output: []interval{{1, 6}},
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.output, s.merge(tc.input))
	}
}
