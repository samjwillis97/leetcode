package main

import "fmt"

func main() {
	nums := []int{1, 2}
	k := 3
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3
	//nums := []int{-1, 100, 3, 99}
	//k := 2
	rotate(nums, k)
	fmt.Println(nums)
}

// Given an array, rotate the array to the right by k steps, where k is non-negative.
//
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

// copy
func rotate(nums []int, k int) {
	if k != 0 {
		toRot := nums[len(nums)-1]
		rarr := append(make([]int, 1), nums[:len(nums)-1]...)
		for i := 0; i < k; i++ {
			if i > 0 {
				toRot = rarr[len(rarr)-1]
				rarr = append(make([]int, 1), rarr[:len(rarr)-1]...)
			}
			rarr[0] = toRot
		}
		copy(nums, rarr)
	}
}
1
//func rotate(nums []int, k int)  {
//
//	n := len(nums)
//	rot_arr := make([]int, n)
//	for i:= 0; i<n; i++ {
//		rot_arr[(i+k)%n] = nums[i]
//	}
//	copy(nums,rot_arr)
//}
