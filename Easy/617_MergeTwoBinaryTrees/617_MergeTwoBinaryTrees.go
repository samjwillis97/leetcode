package mergeTwoBinaryTrees

import (
	"errors"
	"fmt"
)

type NodePair struct {
	l *TreeNode
	r *TreeNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://www.geeksforgeeks.org/merge-two-binary-trees-node-sum/

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// Quick Edge Case Handling
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	if root1.Left == nil && root1.Right == nil {
		head := root2
		head.Val = root2.Val + root1.Val
		return head
	}
	if root2.Left == nil && root2.Right == nil {
		head := root1
		head.Val = root1.Val + root2.Val
		return head
	}
	s := NewStack(5000)
	s.Push(NodePair{
		root1,
		root2,
	})
	for !s.isEmpty() {
		//s.Display()
		n, ok := s.Pop().(NodePair)
		//fmt.Printf("L: %v, R: %v;\n", n.l, n.r)
		if !ok {
			break
		}
		var val int
		fmt.Println(val)
		if n.l != nil {
			val += n.l.Val
		}
		fmt.Println(val)
		if n.r != nil {
			val += n.r.Val
		}
		fmt.Println(val)
		n.l.Val = val
		//n.l.Val += n.r.Val
		//fmt.Printf("New Val: %d\n", n.l.Val)
		if n.l.Left == nil {
			if n.r != nil {
				//fmt.Printf("Append %v to Left of Tree 1\n", n.r.Left)
				n.l.Left = n.r.Left
			}
		} else {
			if n.r != nil {
				//fmt.Println("Push Left Pair")
				s.Push(NodePair{
					n.l.Left,
					n.r.Left,
				})
			}
		}
		if n.l.Right == nil {
			if n.r != nil {
				//fmt.Printf("Append %v to Right of Tree 1\n", n.r.Right)
				n.l.Right = n.r.Right
			}
		} else {
			if n.r != nil {
				//fmt.Println("Push Right Pair")
				s.Push(NodePair{
					n.l.Right,
					n.r.Right,
				})
			}
		}
	}
	// Rest
	return root1
}

func printTree(prefix string, n *TreeNode) {
	if n != nil {
		printTree(prefix+"    ", n.Right)
		fmt.Printf("%s |-- %d\n", prefix, n.Val)
		printTree(prefix+"    ", n.Left)
	}
}

// Stack items stores all the items in the stack
// top stores the index for the top n, -1 on empty, increment on push, decrement on pop
type Stack struct {
	items []interface{}
	top   int
	size  int
}

func NewStack(size int) *Stack {
	return &Stack{
		make([]interface{}, 0),
		-1,
		size,
	}
}

func (s *Stack) Push(a interface{}) error {
	if s.top == s.size-1 {
		return errors.New("stack is full, cannot push")
	}
	s.items = append(s.items, a)
	s.top++
	return nil
}

func (s *Stack) Pop() interface{} {
	if s.top == -1 {
		return nil
	}
	val := s.items[s.top]
	s.items = s.items[:s.top]
	s.top--
	return val
}

func (s *Stack) isEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *Stack) isFull() bool {
	if s.top == s.size-1 {
		return true
	}
	return false
}

func (s *Stack) Peek() interface{} {
	if s.top == -1 {
		return nil
	}
	val := s.items[s.top]
	return val
}

func (s *Stack) Display() {
	fmt.Println("Display")
	for _, val := range s.items {
		fmt.Printf("L: %d R: %d\n", val.(NodePair).l.Val, val.(NodePair).r.Val)
	}
}
