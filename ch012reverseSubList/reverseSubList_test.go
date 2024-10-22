package ch012reversesublist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseSubList(t *testing.T) {
	testCases := []struct {
		head   []int
		p, q   int
		result []int
	}{
		{
			head: []int{1,2,3,4,5},
			p: 2,
			q: 4,
			result: []int{1,4,3,2,5},
		},
		{
			head: []int{1,2,3,4,5,6},
			p: 1,
			q: 4,
			result: []int{4,3,2,1,5,6},
		},
		{
			head: []int{1,2,3,4,5,6,7},
			p: 3,
			q: 6,
			result: []int{1,2,6,5,4,3,7},
		},
	}

	for _, tc := range testCases {
		mainLinkedList := createListNode(tc.head, 0)
		// expectedLinkedList := createListNode(tc.result, 0)
		reversedSubList := reverseSubList(mainLinkedList, tc.p, tc.q)
		expectedSubList := linkedListToList(reversedSubList)
		assert.Equal(t, tc.result, expectedSubList)
	}
}

func TestCreateNodeList(t *testing.T) {
	mainLinkedList := []int{1,2,3,4}
	expectedResult := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, &ListNode{
					4, nil,
				},
			},
		},
	}
	assert.Equal(t, expectedResult, createListNode(mainLinkedList, 0))
}

func TestLinkedListToList(t *testing.T) {
	linkedList := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, &ListNode{
					4, nil,
				},
			},
		},
	}
	expextedList := []int{1,2,3,4}
	assert.Equal(t, expextedList, linkedListToList(linkedList))
}


func createListNode(ls []int, idx int) *ListNode {
	if idx >= len(ls) {
		return nil
	} else {
		return &ListNode{
			Val: ls[idx],
			Next: createListNode(ls, idx + 1),
		}
	}
}

func linkedListToList(head *ListNode) []int {
	var ls = []int{head.Val}
	if head.Next != nil {
		ls = append(ls, linkedListToList(head.Next)...)
	}
	return ls
}