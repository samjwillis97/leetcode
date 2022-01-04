package middleOfLinkedList

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMiddleOfLinkedListExamples(t *testing.T) {
	tests := []struct {
		linkedList string
		want       string
	}{
		{"[1,2,3,4,5]", "[3,4,5]"},
		{"[1,2,3,4,5,6]", "[4,5,6]"},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("MiddleOfLinkedList=%d", i), func(t *testing.T) {
			got := LinkedListToStr(middleNode(ParseLinkedListFromStr(tc.linkedList)))
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
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
