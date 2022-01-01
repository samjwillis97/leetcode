package main

import (
	helpers "LeetCode/Helpers"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node := helpers.ParseIntLinkedListFromStr("1->2->2->1->")
	helpers.PrintIntLinkedListNode(node)
	fmt.Println(isPalindrome234(node))
}

//Input: head = [1,2,2,1]
//Output: true

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome234(head *helpers.IntListNode) bool {
	var numbers = make([]int, 0)
	if head.Next == nil {
		return true
	}
	numbers = AppendValList(head, numbers)

	for i := 0; i < len(numbers)/2; i++ {
		if numbers[i] != numbers[len(numbers)-(i+1)] {
			return false
		}
	}
	return true
}

func AppendValList(head *helpers.IntListNode, list []int) []int {
	list = append(list, head.Val)
	if head.Next != nil {
		list = AppendValList(head.Next, list)
	}
	return list
}
