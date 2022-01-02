package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSumSortedArray(nums, target))
}

// Given a 1-indexed array of integers numbers that is
// already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number.
// Let these two numbers be numbers[index1] and numbers[index2]
// where 1 <= index1 < index2 <= numbers.length.
//
// Return the indices of the two numbers, index1 and index2,
// added by one as an integer array [index1, index2] of length 2.
//
// The tests are generated such that there is exactly one solution.
// You may not use the same element twice.
//
// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9.
// Therefore, index1 = 1, index2 = 2. We return [1, 2].
//

// Two Pointers,
// uses the fact that the array is sorted
// start from both ends, lowest and highest
// for Sum, if the sum is less than the target, shift left pointer to the right
// if sum is greater, shift right pointer left

func twoSumSortedArray(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			break
		}
		if sum < target {
			l += 1
		} else {
			r -= 1
		}
	}
	return []int{l + 1, r + 1}
}
