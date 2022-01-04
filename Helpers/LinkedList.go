package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ParseLinkedListFromStr(input string) *ListNode {
	arr := strings.Split(input, `->`)
	if len(arr) == 0 {
		return nil
	}
	intArr := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		t, err := strconv.Atoi(arr[i])
		if err != nil {
			panic(err)
		}
		intArr[i] = t
	}
	node := &ListNode{intArr[0], nil}
	p := node
	for i := 1; i < len(arr)-1; i++ {
		p.Next = &ListNode{intArr[i], nil}
		p = p.Next
	}
	return node
}

func PrintLinkedListNode(node *ListNode) {
	for p := node; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Println()
}

func ReverseLinkedList(list *ListNode) *ListNode {
	var prev, current, next *ListNode
	prev = nil
	current = list
	next = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
