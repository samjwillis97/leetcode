package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

// Given a string s,
// reverse the order of characters in each word within a sentence
// while still preserving whitespace and initial word order.

// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

func reverseWords(s string) string {
	a := strings.Split(s, " ")
	for ndx, val := range a {
		a[ndx] = reverseString(val)
	}
	return strings.Join(a, " ")
}

func reverseString(s string) string {
	a := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		l := a[i]
		r := a[len(a)-(1+i)]
		a[i] = r
		a[len(a)-(1+i)] = l
	}
	return string(a)
}
