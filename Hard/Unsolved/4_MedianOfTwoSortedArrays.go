// UNSOLVED
package main

import (
	"fmt"
	"math"
)

func main() {
	nums1 := []int{3, 4}
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

//
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
//
// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
//
// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//
// Constraints:
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
//

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// get min/max of nums1
	// check if min/max of nums2 is within nums1
	// use a binary search? to find where it fits in and merge them?

	min1, max1 := minMaxSorted(nums1)
	min2, max2 := minMaxSorted(nums2)

	/*
	 * A [-----]
	 * B            [-----]
	 */
	if max1 < min2 {
		fmt.Println("Nums 1 Lower")
		return findMedianSortedNoOverlap(nums1, nums2)
	} else if max2 < min1 {
		fmt.Println("Nums 2 Lower")
		return findMedianSortedNoOverlap(nums2, nums1)
	}

	/*
	 * A     [-----]
	 * B [---------]
	 */
	return findMedianSortedOverlap(nums1, nums2)
}

func minMaxSorted(nums []int) (int, int) {
	min := nums[0]
	max := nums[len(nums)-1]
	return min, max
}

func findMedianSortedOverlap(nums1 []int, nums2 []int) float64 {
	median1 := medianOfArray(nums1)
	median2 := medianOfArray(nums2)
	if median1 == median2 {
		fmt.Println("Equal Medians")
		return median1
	}

	// Get Medians of Both
	// Only get slice between medians, merge them
	// Find Median
	fmt.Println("Index of Median2 in Nums1")
	fmt.Println(binarySearchSorted(nums1, int(median2)))
	fmt.Println("Index of Median1 in Nums2")
	fmt.Println(binarySearchSorted(nums2, int(median1)))

	return 0
}

func findMedianSortedNoOverlap(lower []int, higher []int) float64 {
	return (float64(lower[len(lower)-1]) + float64(higher[0])) / 2
}

func binarySearchSorted(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1
	case a[mid] > search:
		result, searchCount = binarySearchSorted(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearchSorted(a[mid+1:], search)
		if result >= 0 {
			result += mid + 1
		}
	default:
		result = mid
	}
	searchCount++
	return
}

func medianOfArray(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return float64(nums[0])
	}
	if len(nums)%2 != 0 {
		return float64(nums[len(nums)/2])
	}
	upper := float64(nums[int(math.Ceil(float64(len(nums)/2)))])
	lower := float64(nums[int(math.Ceil(float64(len(nums)/2))-1)])
	return (upper + lower) / 2
}
