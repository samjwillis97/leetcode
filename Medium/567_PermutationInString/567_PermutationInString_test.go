package permutationInString

import (
	"fmt"
	"testing"
)

func TestExamples(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"prosperity", "properties", false},
		{"intention", "execution", false},
		{"rvwrk", "lznomzggwrvrkxecjaq", true},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("Example=%d", i), func(t *testing.T) {
			got := checkInclusion(tc.s1, tc.s2)
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
