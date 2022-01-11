package permutations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("combinations=%d", i), func(t *testing.T) {
			got := permute(tc.nums)
			if !reflect.DeepEqual(got, tc.want) {
				fmt.Println()
				printMatrix(got)
				fmt.Println()
				printMatrix(tc.want)
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
