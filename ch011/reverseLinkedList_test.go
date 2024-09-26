package ch011reverselinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	testCases := []struct {
		input, output *ListNode
	}{
		{input: &ListNode{
			2, &ListNode{
				4, &ListNode{
					6, &ListNode{
						8, &ListNode{
							10, nil,
						},
					},
				},
			},
		}, output: &ListNode{
			10, &ListNode{
				8, &ListNode{
					6, &ListNode{
						4, &ListNode{
							2, nil,
						},
					},
				},
			},
		}},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.output, reverseLinkedList(tc.input))
	}

}
