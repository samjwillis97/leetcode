package main

import (
	"fmt"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}

// Given an integer array nums sorted in non-decreasing order,
// return an array of the squares of each number sorted in non-decreasing order.

func sortedSquares(nums []int) []int {
	a := make([]int, 0, len(nums))
	for _, val := range nums {
		calc := val * val
		if len(a) == 0 {
			a = []int{calc}
		} else {
			ndx := insertLocation(a, calc)
			a = insert(a, calc, ndx)
		}
	}
	return a
}

func insert(a []int, c int, i int) []int {
	return append(a[:i], append([]int{c}, a[i:]...)...)
}

func insertLocation(nums []int, target int) int {
	var m, l, r, val int
	l = 0
	r = len(nums) - 1
	for l <= r {
		m = (r + l) / 2
		val = nums[m]
		if val == target {
			return m
		}
		if val > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
