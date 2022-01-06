package permutationInString

import "fmt"

/*
	Simple
	To Find Permutations of Size 3, Start with Size 2 (n-1)
	1, 2
	2, 1
	Then Insert 3 at all positions
	3, 1, 2
	1, 3, 2
	1, 2, 3 etc.
*/

func checkInclusion(s1 string, s2 string) bool {
	var r int
	var helper func([]byte, int)
	var perms [][]byte

	windowSize := len(s1)
	if windowSize > len(s2) {
		return false
	}
	if windowSize == len(s2) {
		map1 := toMap([]byte(s1))
		map2 := toMap([]byte(s2))
		if len(map1) != len(map2) {
			return false
		}
		for key := range map1 {
			if _, ok := map2[key]; !ok {
				return false
			}
		}
		return true
	}
	for l := 0; r < len(s2); l++ {
		perms = [][]byte{}
		r = l + windowSize
		fmt.Printf("L: %d, R: %d, W: %s\n", l, r, s2[l:r])
		if s1 == s2[l:r] {
			return true
		}
		helper = func(a []byte, n int) {
			if n == 1 {
				tmp := make([]byte, len(a))
				copy(tmp, a)
				perms = append(perms, tmp)
			} else {
				for i := 0; i < n; i++ {
					helper(a, n-1)
					if n%2 != 0 {
						a = swap(a, 0, len(a)-1)
					} else {
						a = swap(a, i, len(a)-1)
					}
				}
			}
		}
		helper([]byte(s2[l:r]), windowSize)
		for _, perm := range perms {
			if l == 8 {
				fmt.Println(string(perm))
			}
			if s1 == string(perm) {
				return true
			}
		}

	}
	return false
}

/*
	Heaps
	1.	The algorithm generates (n-1)! permutations of the first n-1 elements,
		adjoining the last element to each of these. This will generate all of
		the permutations that end with the last element.
    2.	If n is odd, swap the first and last element and if n is even, then
		swap the ith element (i is the counter starting from 0) and the last
		element and repeat the above algorithm till i is less than n.
    3.	In each iteration, the algorithm will produce all the permutations
		that end with the current last element.
*/

func swap(a []byte, i, j int) []byte {
	a[i], a[j] = a[j], a[i]
	return a
}

func toMap(a []byte) map[byte]struct{} {
	charMap := map[byte]struct{}{}
	for _, val := range a {
		charMap[val] = struct{}{}
	}
	return charMap
}
