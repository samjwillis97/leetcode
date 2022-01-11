package combinations

import "fmt"

// Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].
// You may return the answer in any order.

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	// Backtrack + DFS
	// Initially
	// Start = 1
	// curComb, currentCombination, init empty []int
	var comb func(start int, curComb []int)
	comb = func(start int, curComb []int) {
		fmt.Println(curComb)

		// Stop Condition
		if len(curComb) == k {
			// Make a Copy of Current Combination
			// Append to Results
			dst := make([]int, k)
			copy(dst, curComb)
			result = append(result, dst)
		}

		// General Case
		// Iterates from 1 -> n
		for i := start; i <= n; i++ {
			// FOR n = 4, k = 2
			// First Ever Call [start], should be [1]
			curComb = append(curComb, i)
			// calls comb(2, curComb)
			// will go through until
			comb(i+1, curComb)
			// will first return [1, 2]
			curComb = curComb[:len(curComb)-1]
			// curComb = [1]
		}
		return
	}
	comb(1, make([]int, 0))
	return result
}
