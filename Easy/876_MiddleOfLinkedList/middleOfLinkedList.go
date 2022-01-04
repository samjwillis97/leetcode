package middleOfLinkedList

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a singly linked list, return the middle node of the linked list.
//
// If there are two middle nodes, return the second middle node.

func middleNode(head *ListNode) *ListNode {
	length := 0
	for l := head; l != nil; l = l.Next {
		length++
	}
	count := int(math.Ceil(float64(length+1) / 2))
	mid := head
	for i := 1; i < count; i++ {
		mid = mid.Next
	}

	return mid
}
