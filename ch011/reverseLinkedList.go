package ch011reverselinkedlist

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	var previous *ListNode = nil
	current := head
	var next *ListNode = current.Next
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	return previous
}