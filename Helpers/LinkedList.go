package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

type IntListNode struct {
	Val  int
	Next *IntListNode
}

func ParseIntLinkedListFromStr(input string) *IntListNode {
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
	node := &IntListNode{intArr[0], nil}
	p := node
	for i := 1; i < len(arr)-1; i++ {
		p.Next = &IntListNode{intArr[i], nil}
		p = p.Next
	}
	return node
}

func PrintIntLinkedListNode(node *IntListNode) {
	for p := node; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Println()
}

func ReverseIntLinkedList(list *IntListNode) *IntListNode {
	var prev, current, next *IntListNode
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
