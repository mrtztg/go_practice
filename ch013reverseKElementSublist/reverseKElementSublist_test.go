package ch013reversekelementsublist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		head   []int
		k      int
		result []int
	}{
		{
			head:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:      3,
			result: []int{3, 2, 1, 6, 5, 4, 8, 7},
		},
		{
			head: []int{1, 2, 3, 4, 5},
			k: 2,
			result: []int{2, 1, 4, 3, 5},
		},
		{
			head: []int{1, 2, 3, 4, 5, 6, 7},
			k: 4,
			result: []int{4, 3, 2, 1, 7, 6, 5},
		},
		{
			head: []int{100, 200, 300, 400, 500},
			k: 5,
			result: []int{500, 400, 300, 200, 100},
		},
	}

	for _, tc := range testCases {
		mainLinkedList := createListNode(tc.head, 0)
		// expectedLinkedList := createListNode(tc.result, 0)
		reversedSubList := Solution(mainLinkedList, tc.k)
		expectedSubList := linkedListToList(reversedSubList)
		assert.Equal(t, tc.result, expectedSubList)
	}
}

func TestCreateNodeList(t *testing.T) {
	mainLinkedList := []int{1, 2, 3, 4}
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
	expextedList := []int{1, 2, 3, 4}
	assert.Equal(t, expextedList, linkedListToList(linkedList))
}

func createListNode(ls []int, idx int) *ListNode {
	if idx >= len(ls) {
		return nil
	} else {
		return &ListNode{
			Val:  ls[idx],
			Next: createListNode(ls, idx+1),
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
