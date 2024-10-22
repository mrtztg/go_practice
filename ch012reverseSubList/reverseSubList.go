package ch012reversesublist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseSubList(head *ListNode, p, q int) *ListNode {
	if p == q {
		return head
	}
	var previous *ListNode
	var cursor = head
	i := 1
	for ; i < p; i++ {
		previous = cursor
		cursor = cursor.Next
	}
	lastUntouched := previous

	var next *ListNode
	var reversedListTail = cursor
	for ; cursor != nil && i <= q; i++ {
		next = cursor.Next
		cursor.Next = previous
		previous = cursor
		cursor = next
	}
	if lastUntouched != nil {
		lastUntouched.Next = previous
	} else {
		head = previous
	}
	reversedListTail.Next = cursor
	return head
}
