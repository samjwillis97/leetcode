package matrix01

import (
	"fmt"
)

// Given an m x n binary matrix mat,
// return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.

func updateMatrix(mat [][]int) [][]int {
	// For Each Square that != 0
	// Add All Neighbours to a Stack, Set Counter = 0
	// If Stack is not empty
	// if any neighbour == 0
	// 		space = 1
	// else add neighbour to a stack
	//
	m := len(mat)
	if m == 0 {
		return mat
	}

	n := len(mat[0])
	if n == 0 {
		return mat
	}
	fmt.Printf("M: %d, N: %d\n", m, n)

	tempQ := &Queue{top: -1}
	mainQ := &Queue{top: -1}

	for i, row := range mat {
		for j, col := range row {
			if col != 0 {
				//fmt.Printf("I: %d, J: %d, not 0\n", i, j)
				dist := 1
				if i > 0 {
					// Up
					if mat[i-1][j] == 0 {
						continue
					} else {
						mainQ.Push(Node{i - 1, j})
					}
				}
				if i < m-1 {
					// Down
					if mat[i+1][j] == 0 {
						continue
					} else {
						mainQ.Push(Node{i + 1, j})
					}
				}
				if j > 0 {
					// Left
					if mat[i][j-1] == 0 {
						continue
					} else {
						mainQ.Push(Node{i, j - 1})
					}
				}
				if j < n-1 {
					// Left
					if mat[i][j+1] == 0 {
						continue
					} else {
						mainQ.Push(Node{i, j + 1})
					}
				}
				for !mainQ.isEmpty() {
					dist++
					el, ok := mainQ.Pop().(Node)
					if i == 9 && j == 0 {
						fmt.Printf("Currently I: %d, J: %d\n", el.i, el.j)
					}
					if !ok {
						fmt.Println("Error TypeCasting from Main Queue")
						return nil
					}
					if el.i > 0 {
						if mat[el.i-1][el.j] == 0 {
							break
						} else {
							tempQ.Push(Node{el.i - 1, el.j})
						}
					}
					if el.i < m-1 {
						if mat[el.i+1][el.j] == 0 {
							break
						} else {
							tempQ.Push(Node{el.i + 1, el.j})
						}
					}
					if el.j > 0 {
						if mat[el.i][el.j-1] == 0 {
							break
						} else {
							tempQ.Push(Node{el.i, el.j - 1})
						}
					}
					if el.j < n {
						if mat[el.i][el.j+1] == 0 {
							break
						} else {
							tempQ.Push(Node{el.i, el.j + 1})
						}
					}
					if mainQ.isEmpty() {
						mainQ = tempQ
					}
				}
				mat[i][j] = dist
			}
		}
	}

	return mat
}

type Node struct {
	i int
	j int
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
