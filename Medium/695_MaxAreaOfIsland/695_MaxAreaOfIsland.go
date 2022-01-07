package maxAreaOfIsland

import (
	"fmt"
	"strconv"
)

// You are given an m x n binary matrix grid.
// An island is a group of 1's (representing land)
// connected 4-directionally (horizontal or vertical.)
// You may assume all four edges of the grid are surrounded by water.
//
// The area of an island is the number of cells with a value 1 on the island.
//
// Return the maximum area of an island in grid. If there is no island, return 0.

func maxAreaOfIsland(grid [][]int) int {
	// TODO: empty case
	printGrid(grid)
	largest := 0
	seen := map[string]struct{}{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && !beenSeen(i, j, seen) {
				size, newSeen := matchingArea(grid, i, j, 1, "", seen)
				seen = newSeen
				//fmt.Println(size)
				if size > largest {
					largest = size
				}
			}
		}

	}
	//size, _ := matchingArea(grid, 3, 8, 1, "", seen)
	return largest
}

func printGrid(a [][]int) {
	printL := len(a[0])
	fmt.Printf(" |")
	for i := 0; i < printL; i++ {
		fmt.Printf("%d|", i)
	}
	fmt.Printf("\n")
	for i := 0; i < printL; i++ {
		fmt.Printf("-|")
	}
	fmt.Printf("\n")

	for ndx, row := range a {
		fmt.Printf("%d", ndx)
		for _, i := range row {
			fmt.Printf("|%d", i)
		}
		fmt.Printf("|\n")
	}
}

func matchingArea(a [][]int, i int, j int, match int, from string, seen map[string]struct{}) (int, map[string]struct{}) {
	//fmt.Printf("i: %d, j: %d\n", i, j)
	size := 1
	//printImage(image)
	iBound := len(a) - 1
	jBound := len(a[0]) - 1
	if i < iBound {
		iNew := i + 1
		if a[iNew][j] == match && !beenSeen(iNew, j, seen) && from != "d" {
			//fmt.Println("V")
			add, newSeen := matchingArea(a, iNew, j, match, "u", seen)
			seen = newSeen
			size += add
		} else {
			//fmt.Printf("Cannot Go Down To: %d, %d\n", iNew, j)
		}
	}
	if i > 0 {
		iNew := i - 1
		if a[iNew][j] == match && !beenSeen(iNew, j, seen) && from != "u" {
			//fmt.Println("^")
			add, newSeen := matchingArea(a, iNew, j, match, "d", seen)
			seen = newSeen
			size += add
		} else {
			//fmt.Printf("Cannot Go Up To: %d, %d\n", iNew, j)
		}
	}
	if j < jBound {
		jNew := j + 1
		if a[i][jNew] == match && !beenSeen(i, jNew, seen) && from != "r" {
			//fmt.Println("->")
			add, newSeen := matchingArea(a, i, jNew, match, "l", seen)
			seen = newSeen
			size += add
		} else {
			//fmt.Printf("Cannot Go Right To: %d, %d\n", i, jNew)
		}
	}
	if j > 0 {
		jNew := j - 1
		if a[i][jNew] == match && !beenSeen(i, jNew, seen) && from != "l" {
			//fmt.Println("<-")
			add, newSeen := matchingArea(a, i, jNew, match, "r", seen)
			seen = newSeen
			size += add
		} else {
			//fmt.Printf("Cannot Go Left To: %d, %d\n", i, jNew)
		}
	}
	return size, seen
}

func beenSeen(i, j int, seen map[string]struct{}) bool {
	str := strconv.Itoa(i) + "-" + strconv.Itoa(j)
	//fmt.Printf("Checking: %d, %d\n", i, j)
	if _, ok := seen[str]; !ok {
		// Not Seen Row
		seen[str] = struct{}{}
		return false
	}
	//fmt.Println("Seen")
	return true
}
