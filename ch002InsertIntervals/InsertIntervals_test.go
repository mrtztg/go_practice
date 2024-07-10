package ch002InsertIntervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolutionInsertInterval(t *testing.T) {
	testCases := []struct {
		intervals   []Interval
		newInterval Interval
		output      []Interval
	}{
		{
			[]Interval{{1, 3}, {5, 7}, {8, 12}},
			Interval{4, 6},
			[]Interval{{1, 3}, {4, 7}, {8, 12}},
		}, {
			[]Interval{{1, 3}, {5, 7}, {8, 12}},
			Interval{4, 10},
			[]Interval{{1, 3}, {4, 12}},
		}, {
			[]Interval{{2, 3}, {5, 7}},
			Interval{1, 4},
			[]Interval{{1, 4}, {5, 7}},
		},
	}
	s := &solution{}
	for _, tc := range testCases {
		assert.Equal(t, tc.output, s.insert(tc.intervals, tc.newInterval))
	}
}
