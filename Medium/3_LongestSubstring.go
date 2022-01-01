package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

//Given a string s, find the length of the longest substring without repeating characters.
//
//Example 1:
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.

// NEED TO KEEP INDEX OF WHERE EACH WAS FOUND MAP
// abcabcbb
// 01234567
// 012
// 123
// 23

//func lengthOfLongestSubstring(s string) int {
//	if len(s) == 0 {
//		return 0
//	}
//	hmap := map[uint8]int{s[0]: 0}
//	maxLen := 1
//	for i, j := 0, 1; j < len(s); j++ {
//		if v, ok := hmap[s[j]]; ok && v >= i {
//			i = v + 1
//		} else {
//			if maxLen < j-i+1 {
//				maxLen = j - i + 1
//			}
//		}
//		hmap[s[j]] = j
//	}
//	return maxLen
//}

func lengthOfLongestSubstring(s string) int {
	var running int
	seen := map[uint8]int{}
	length := 0

	if len(s) <= 1 {
		return len(s)
	}

	for pos := 0; pos < len(s); pos++ {
		val, ok := seen[s[pos]]
		if !ok {
			running++
		} else {
			pos = val + 1
			seen = map[uint8]int{}
			running = 1
		}
		if running > length {
			length = running
		}
		seen[s[pos]] = pos
	}
	if running > length {
		length = running
	}

	return length
}
