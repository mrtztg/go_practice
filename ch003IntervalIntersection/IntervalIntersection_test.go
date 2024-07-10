package ch003IntervalIntersection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerge(t *testing.T) {
	s := solution{}

	testcases := []struct {
		arr1, arr2, output []Interval
	}{
		{
			arr1: []Interval{
				{1, 3},
				{5, 6},
				{7, 9},
			},
			arr2: []Interval{
				{2, 3},
				{5, 7},
			},
			output: []Interval{
				{2, 3},
				{5, 6},
				{7, 7},
			},
		}, {
			arr1: []Interval{
				{1, 3},
				{5, 7},
				{9, 12},
			},
			arr2: []Interval{
				{5, 10},
			},
			output: []Interval{
				{5, 7},
				{9, 10},
			},
		},
	}
	for _, tc := range testcases {
		assert.Equal(t, tc.output, s.merge(tc.arr1, tc.arr2))
	}
}
