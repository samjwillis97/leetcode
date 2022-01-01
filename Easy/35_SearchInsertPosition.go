package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}

// Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
// Therefore -> Binary Search
//
//Example 1:
//
//Input: nums = [1,3,5,6], target = 5
//Output: 2

func searchInsert(nums []int, target int) int {
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
	if target > val {
		return m + 1
	} else {
		return m
	}
}
