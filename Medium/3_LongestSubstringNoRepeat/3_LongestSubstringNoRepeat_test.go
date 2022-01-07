package longestSubstringNoRepeat

import (
	"fmt"
	"testing"
)

func TestLongestSubstringNoRepeat(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("LengthOfLongestSubstring=%d", i), func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
