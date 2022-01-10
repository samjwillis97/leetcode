package rottingOranges

import (
	"fmt"
	"strconv"
)

//You are given an m x n grid where each cell can have one of three values:
//
//0 representing an empty cell,
//1 representing a fresh orange, or
//2 representing a rotten orange.
//
//Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
//
//Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

func orangesRotting(grid [][]int) int {

	rows := len(grid)
	if rows == 0 {
		return -1
	}
	cols := len(grid[0])
	if cols == 0 {
		return -1
	}

	time := 0
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
	}

	q := &Queue{top: -1}
	fresh := map[string]struct{}{}

	for row, rowVal := range grid {
		for col, colVal := range rowVal {
			if colVal == 0 {
				dist[row][col] = -1
			} else if colVal == 1 {
				dist[row][col] = 9223372036854775807
				fresh[strconv.Itoa(row)+"-"+strconv.Itoa(col)] = struct{}{}
			} else if colVal == 2 {
				q.Push(Node{row, col})
				dist[row][col] = 0
			}
		}
	}

	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for !q.isEmpty() {
		current, ok := q.Pop().(Node)
		if !ok {
			fmt.Println("Error TypeCasting from Main Queue")
			return -1
		}
		for _, val := range dir {
			new_r := current.i + val[0]
			new_c := current.j + val[1]
			if new_r >= 0 && new_c >= 0 && new_r < rows && new_c < cols {
				// If the Distance/Time of the current square
				// is greater swuare stepping from add 1
				// will be greater if not visited due to max int default
				// then enqueue
				if dist[new_r][new_c] > dist[current.i][current.j]+1 {
					delete(fresh, strconv.Itoa(new_r)+"-"+strconv.Itoa(new_c))
					dist[new_r][new_c] = dist[current.i][current.j] + 1
					q.Push(Node{new_r, new_c})
					if dist[new_r][new_c] > time {
						time = dist[new_r][new_c]
					}
				}
			}
		}
	}
	if len(fresh) > 0 {
		return -1
	}
	return time
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
