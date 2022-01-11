package mergeTwoLists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return nil
	}
	// Compare Head of Both Lists
	// Make smaller head the current element
	// Recursively Solve with parameters
	// 		next node of the smaller head
	//		the other node
	// Function will return the smaller with the rest linked
	var current *ListNode
	var other *ListNode
	if list1.Val < list2.Val {
		current = list1
		other = list2
	} else {
		current = list2
		other = list1
	}

	current.Next = mergeTwoLists(current.Next, other)

	// My First Attempt
	// Find Where Current List 2 Node Fits into List 1
	// if List2.Next != nil Do Again...
	// Need to make sure not to chop bits of the list off
	//if list2.Val > list1.Val {
	//	// List 2 Head is between After List 1 Head
	//	if list1.Next == nil {
	//		// List 1 Ended, Append List 2
	//		list1.Next = list2
	//		return list1
	//	}
	//	if list1.Next.Val < list2.Val {
	//		// List 2 Fits Between List 1 Head and Next
	//		// List1.Next.Val
	//		// Let List 1 Next = List 2 Head
	//	}
	//} else {
	//
	//}
	return current
}

//func ReverseLinkedList(list *ListNode) *ListNode {
//	var prev, current, next *ListNode
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
