package main

import (
	"LeetCode/Helpers"
)

func main() {
	l1 := helpers.ParseIntLinkedListFromStr("0->")
	l2 := helpers.ParseIntLinkedListFromStr("7->3->")
	helpers.PrintIntLinkedListNode(l1)
	helpers.PrintIntLinkedListNode(l2)
	helpers.PrintIntLinkedListNode(addTwoNumbers(l1, l2))
}

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//Input: l1 = [2,4,3], l2 = [5,6,4]
//Output: [7,0,8]
//Explanation: 342 + 465 = 807.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *helpers.IntListNode, l2 *helpers.IntListNode) *helpers.IntListNode {
	var head, prev *helpers.IntListNode
	prev = nil
	carry := 0
	for i := 0; l1 != nil || l2 != nil || carry != 0; i++ {
		newVal := 0
		if l1 != nil {
			newVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			newVal += l2.Val
			l2 = l2.Next
		}
		newVal += carry
		if newVal > 9 {
			newVal = newVal - 10
			carry = 1
		} else {
			carry = 0
		}
		sum := helpers.IntListNode{
			Val:  newVal,
			Next: nil,
		}
		if i == 0 {
			head = &sum
		}
		if prev != nil {
			prev.Next = &sum
		}
		prev = &sum
	}
	return head
}

//func reverseIntLinkedList(list *helpers.IntListNode) *helpers.IntListNode {
//	var prev, current, next *helpers.IntListNode
//	prev = nil
//	current = list
//	next = nil
//
//	for current != nil {
//		next = current.Next
//		current.Next = prev
//		prev = current
//		current = next
//	}
//	return prev
//}

//1) Divide the list in two parts - first node and
//rest of the linked list.
//2) Call reverse for the rest of the linked list.
//3) Link rest to first.
//4) Fix head pointer

//Initialize three pointers prev as NULL, curr as head and next as NULL.
//Iterate through the linked list. In loop, do following.
//// Before changing next of current,
//// store next node
//next = curr->next
//// Now change next of current
//// This is where actual reversing happens
//curr->next = prev
//// Move prev and curr one step forward
//prev = curr
//curr = next
