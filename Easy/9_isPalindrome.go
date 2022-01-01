package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome9(0))
}

// Given an integer x, return true if x is palindrome integer.
//
// An integer is a palindrome when it reads the same backward as forward.
//
// For example, 121 is a palindrome while 123 is not.
//
// Constraints
// -231 <= x <= 231 - 1

func isPalindrome9(x int) bool {
	xStr := strconv.Itoa(x)
	length := len(xStr)
	palindrome := false
	if length == 1 {
		return true
	}
	for i := 0; i < length/2; i++ {
		if xStr[i] == xStr[length-(i+1)] {
			palindrome = true
		} else {
			palindrome = false
			break
		}
	}
	return palindrome
}
