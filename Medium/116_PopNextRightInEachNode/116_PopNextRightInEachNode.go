package popNextRightInEachNode

import (
	"fmt"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	mainQ := &Queue{top: -1}
	tempQ := &Queue{top: -1}

	mainQ.Push(root)

	if !mainQ.isEmpty() {
		n, ok := mainQ.Pop().(*Node)
		if !ok {
			fmt.Println("Error TypeCasting from Main Queue")
			return nil
		}

		next, ok := mainQ.Peek().(*Node)
		if !ok {
			fmt.Printf("Pointing %d Next -> nil", n.Val)
			n.Next = nil
		} else {
			fmt.Printf("Pointing %d Next -> %d", next.Val)
			n.Next = next
		}

		if n.Left != nil {
			tempQ.Push(n.Left)
		}
		if n.Right != nil {
			tempQ.Push(n.Right)
		}
		if mainQ.isEmpty() {
			mainQ = tempQ
		}
	}
	return root
}

// Queue FIFO
type Queue struct {
	items []interface{}
	top   int
}

func (s *Queue) Push(a interface{}) error {
	s.items = append(s.items, a)
	s.top++
	return nil
}

func (s *Queue) Pop() interface{} {
	if s.top == -1 {
		return nil
	}
	val := s.items[0]
	s.items = s.items[1:]
	s.top--
	return val
}

func (s *Queue) isEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *Queue) Peek() interface{} {
	if s.top == -1 {
		return nil
	}
	val := s.items[0]
	return val
}

func (s *Queue) Length() int {
	return s.top + 1
}

func (s *Queue) Display() {
	for _, val := range s.items {
		fmt.Println(val)
	}
}
