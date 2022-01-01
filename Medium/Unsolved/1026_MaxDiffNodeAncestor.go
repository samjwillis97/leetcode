// UNSOLVED
package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := ParseTreeFromInput(`
		[8,3,10,1,6,null,14,null,null,4,7,13]
	`)
	PrintTree("", tree)
	maxAncestorDiff(tree)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	var results int
	walk(root, root.Val, root.Val, &results)
	return 0
}

func walk(root *TreeNode, min int, max int, results *int) {
	if root == nil {
		return
	}

}

// ONLY FOR TESTING

func ParseTreeFromInput(inputStr string) *TreeNode {
	var input []*int
	if err := json.Unmarshal([]byte(inputStr), &input); err != nil {
		panic(err)
	}
	if len(input) == 0 {
		return nil
	} else if len(input) == 1 {
		return &TreeNode{*input[0], nil, nil}
	}
	r := &TreeNode{*input[0], nil, nil}
	input = input[1:]
	stack := []*TreeNode{r}
	stack1 := []*TreeNode{}
	for len(stack) != 0 {
		p := stack[0]
		stack = stack[1:]
		if len(input) != 0 {
			i := input[0]
			input = input[1:]
			if i != nil {
				p.Left = &TreeNode{*i, nil, nil}
				stack1 = append(stack1, p.Left)
			}
		}
		if len(input) != 0 {
			i := input[0]
			input = input[1:]
			if i != nil {
				p.Right = &TreeNode{*i, nil, nil}
				stack1 = append(stack1, p.Right)
			}
		}
		if len(stack) == 0 {
			stack = stack1
			stack1 = []*TreeNode{}
		}
	}
	return r
}

func PrintTree(prefix string, n *TreeNode) {
	if n != nil {
		PrintTree(prefix+"    ", n.Right)
		fmt.Printf("%s |-- %d\n", prefix, n.Val)
		PrintTree(prefix+"    ", n.Left)
	}
}
