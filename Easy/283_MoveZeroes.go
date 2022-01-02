package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}

// Given an integer array nums,
// move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
// Constraints:
// 1 <= nums.length <= 104
// -2^31 <= nums[i] <= 2^31 - 1
//
// Example
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Two Pointer, one for Iterating Array, other for managing Zero's?

func moveZeroes(nums []int) {
	arrNdx := 0
	zerNdx := len(nums)
	for arrNdx < zerNdx {
		fmt.Println(nums)
		val := nums[arrNdx]
		if val == 0 {
			if arrNdx == 0 {
				nums = append(nums[1:], 0)
			} else {
				nums = append(nums[:arrNdx], nums[arrNdx+1:]...)
				nums = append(nums, 0)
			}
			zerNdx -= 1
			arrNdx = arrNdx - 1
		}
		arrNdx++
	}
	fmt.Println(nums)
}
