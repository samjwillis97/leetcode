package matrix01

import (
	"fmt"
	"math"
)

// Given an m x n binary matrix mat,
// return the distance of the nearest 0 for each cell.
// The distance between two adjacent cells is 1.

func oldUpdateMatrix(mat [][]int) [][]int {
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

//// C++ Solution
// Reverse of Old Solution
// Start BFS From a 0 Element Update all 1's in its path..
// Algorithm:
// 	- For our BFS routine, we keep a queue, q to maintain the queue of cells to be examined next.
//  - We start by adding all the cells with 0s to q.
//  - Intially, distance for each 0 cell is 0 and distance for each 1 is INT_MAX, which is updated during the BFS.
//  - Pop the cell from queue, and examine its neighbors.
// 		If the new calculated distance for neighbor {i,j} is smaller,
//		we add {i,j} to q and update dist[i][j].

func updateMatrix(mat [][]int) [][]int {
	// Edge Cases
	rows := len(mat)
	if rows == 0 {
		return mat
	}
	cols := len(mat[0])
	if cols == 0 {
		return mat
	}

	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
	}

	q := &Queue{top: -1}

	for row, rowVal := range mat {
		for col, colVal := range rowVal {
			if colVal == 0 {
				dist[row][col] = 0
				q.Push(Node{row, col})
			} else {
				dist[row][col] = math.MaxInt
			}
		}
	}
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for !q.isEmpty() {
		current, ok := q.Pop().(Node)
		if !ok {
			fmt.Println("Error TypeCasting from Main Queue")
			return nil
		}
		for _, val := range dir {
			new_r := current.i + val[0]
			new_c := current.j + val[1]
			if new_r >= 0 && new_c >= 0 && new_r < rows && new_c < cols {
				if dist[new_r][new_c] > dist[current.i][current.j]+1 {
					dist[new_r][new_c] = dist[current.i][current.j] + 1
					q.Push(Node{new_r, new_c})
				}
			}
		}
	}

	// Put All 0's into Q
	return dist
}

//class Solution {
//public:
//	vector<vector<int>> updateMatrix(vector<vector<int>>& matrix) {
//		int rows = matrix.size();
//		if (rows == 0)
//			return matrix;
//		int cols = matrix[0].size();
//		vector<vector<int>> dist(rows, vector<int> (cols, INT_MAX));
//		queue<pair<int, int>> q;
//		for (int i = 0; i < rows; i++) {
//			for (int j = 0; j < cols; j++) {
//				if (matrix[i][j] == 0) {
//					dist[i][j] = 0;
//					q.push({ i, j }); //Put all 0s in the queue.
//				}
//			}
//		}
//
//		int dir[4][2] = { { -1, 0 }, { 1, 0 }, { 0, -1 }, { 0, 1 } };
//		while (!q.empty()) {
//			pair<int, int> curr = q.front();
//			q.pop();
//			for (int i = 0; i < 4; i++) {
//				int new_r = curr.first + dir[i][0], new_c = curr.second + dir[i][1];
//				if (new_r >= 0 && new_c >= 0 && new_r < rows && new_c < cols) {
//					if (dist[new_r][new_c] > dist[curr.first][curr.second] + 1) {
//						dist[new_r][new_c] = dist[curr.first][curr.second] + 1;
//						q.push({ new_r, new_c });
//					}
//				}
//			}
//		}
//		return dist;
//	}
//};

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
