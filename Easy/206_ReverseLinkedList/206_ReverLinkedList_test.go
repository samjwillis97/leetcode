package reverseLinkedList

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list string
		want string
	}{
		{"[1,2,4]", "[4,2,1]"},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("TestMergeTwoLists=%d", i), func(t *testing.T) {
			got := reverseList(ParseLinkedListFromStr(tc.list))
			if got != ParseLinkedListFromStr(tc.want) {
				fmt.Println("Got")
				PrintLinkedListNode(got)
				fmt.Println("Want")
				PrintLinkedListNode(ParseLinkedListFromStr(tc.want))

				t.Fatalf("Lists Mismatched")
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

func PrintLinkedListNode(node *ListNode) {
	for p := node; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Println()
}
