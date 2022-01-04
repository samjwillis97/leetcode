package removeNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	for l := head; l != nil; l = l.Next {
		length++
	}

	ndx := length - n
	i := 0

	if ndx == 0 {
		return head.Next
	}

	var prev *ListNode
	for p := head; p != nil; p = p.Next {
		if i == ndx {
			prev.Next = p.Next
		}
		prev = p
		i++
	}

	return head
}
