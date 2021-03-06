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
	arr := strings.Split(input[1:len(input)-1], `,`)
	if len(arr) == 0 {
		return nil
	}
	intArr := make([]int, len(arr))
	for ndx, val := range arr {
		t, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		intArr[ndx] = t
	}
	node := &ListNode{intArr[0], nil}
	p := node
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{intArr[i], nil}
		p = p.Next
	}
	return node
}

func LinkedListToStr(node *ListNode) string {
	var str strings.Builder
	i := 0
	str.WriteString("[")
	for p := node; p != nil; p = p.Next {
		if i != 0 {
			str.WriteString(",")
		}
		str.WriteString(fmt.Sprintf("%d", p.Val))
		i++
	}
	str.WriteString("]")
	return str.String()
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
