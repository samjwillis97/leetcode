package mergeTwoBinaryTrees

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMergeTwoBinaryTree(t *testing.T) {
	tests := []struct {
		tree1 string
		tree2 string
		want  string
	}{
		{"[1,3,2,5]", "[2,1,3,null,4,null,7]", "[3,4,5,5,4,null,7]"},
		{"[1,2,null,3]", "[1,null,2,null,3]", "[2,2,2,3,null,null,3]"},
		{"[4,-9,5,null,-1,null,8,-6,0,7,null,null,-2,null,null,null,null,-3]", "[5]", "[9,-9,5,null,-1,null,8,-6,0,7,null,null,-2,null,null,null,null,-3]"},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("TestMergeTwoBinaryTree=%d", i), func(t *testing.T) {
			got := mergeTrees(ParseTreeFromInput(tc.tree1), ParseTreeFromInput(tc.tree2))
			if got != ParseTreeFromInput(tc.want) {
				fmt.Println("Got")
				printTree("", got)
				fmt.Println("Want")
				printTree("", ParseTreeFromInput(tc.want))

				t.Fatalf("Tree Mismatched")
			} else {
				t.Logf("Success !")
			}
		})
	}
}

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
