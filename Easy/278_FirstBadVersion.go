package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(3))
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	var m int
	var curr bool
	l := 1
	r := n
	for l <= r {
		m = (r + l) / 2
		curr = isBadVersion(m)
		fmt.Printf("L: %d, R: %d, M: %d\n", l, r, m)
		if curr {
			//fmt.Printf("Bad Version\n")
			r = m - 1
		} else {
			//fmt.Printf("Good Version\n")
			l = m + 1
		}
	}
	fmt.Printf("L: %d, R: %d, M: %d\n", l, r, m)
	if curr {
		return m
	} else {
		return m + 1
	}
}

func isBadVersion(n int) bool {
	fmt.Println("Called with:", n)
	return n >= 2
}
