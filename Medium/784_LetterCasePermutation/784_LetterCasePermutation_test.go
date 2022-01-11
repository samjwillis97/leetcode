package letterCasePermutation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLetterPerms(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("letterCasePermutation=%d", i), func(t *testing.T) {
			got := letterCasePermutation(tc.s)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("\ngot  %v;\nwant %v;", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func printMatrix(a [][]int) {
	for _, row := range a {
		for _, i := range row {
			fmt.Printf("|%d", i)
		}
		fmt.Printf("|\n")
	}
}
