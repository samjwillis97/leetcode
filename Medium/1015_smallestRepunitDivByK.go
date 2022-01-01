package main

import (
	"fmt"
)

func main() {
	fmt.Println(smallestRepunitDivByK(49993))
}

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	if k == 1 {
		return 1
	}
	length := 1
	for mod := 1; mod != 0; length++ {
		mod = (mod*10 + 1) % k
		if mod == 0 {
			break
		}
	}
	return length + 1
}
