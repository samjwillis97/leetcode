package reverseLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var current, prev, next *ListNode
	current = head
	prev = nil
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
