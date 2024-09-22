package ch010findcorruptpair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCorruptPair(t *testing.T) {
	testCases := []struct{
		input, output []int
	} {
		{[]int{3,1,2,5,2}, []int{2,4}},
		{[]int{3,1,2,3,6,4}, []int{3,5}},
	}
	for _, tc := range testCases {
		assert.ElementsMatch(t, tc.output, findCorruptPair(tc.input))
	}
}