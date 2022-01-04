package main

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseBytes(s)
}

//Write a function that reverses a string. The input string is given as an array of characters s.
//You must do this by modifying the input array in-place with O(1) extra memory.

func reverseBytes(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		l := s[i]
		r := s[len(s)-(1+i)]
		s[i] = r
		s[len(s)-(1+i)] = l
	}
}

// Can Be done in a single line
// s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
