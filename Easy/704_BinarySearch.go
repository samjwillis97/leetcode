package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 2
	fmt.Println(binarySearch(nums, target))
}

// Given an array of integers nums which is sorted in ascending order,
// and an integer target, write a function to search target in nums.
// If target exists, then return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.

//Input: nums = [-1,0,3,5,9,12], target = 9
//Output: 4

func binarySearch(nums []int, target int) (result int) {
	lower := 0
	upper := len(nums)
	for lower < upper {
		mid := (upper + lower) / 2
		if nums[mid] == target {
			return mid
		}
		// midpoint is less than search target
		// take upper half
		if nums[mid] < target {
			lower = mid + 1
		} else {
			upper = mid
		}
	}
	return -1
}

func recursiveBinarySearch(nums []int, target int) (result int) {
	mid := len(nums) / 2
	val := nums[mid]

	if len(nums) == 0 {
		result = -1
	}
	if val > target {
		result = recursiveBinarySearch(nums[:mid], target)
	} else if val < target {
		result = recursiveBinarySearch(nums[mid+1:], target)
		if result >= 0 {
			result += mid + 1
		}
	} else {
		result = mid
	}
	return
}
