package combinations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinations(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want [][]int
	}{
		{4, 2, [][]int{
			{2, 4},
			{3, 4},
			{2, 3},
			{1, 2},
			{1, 3},
			{1, 4},
		}},
		{4, 3, [][]int{
			{},
		}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("combinations=%d", i), func(t *testing.T) {
			got := combine(tc.n, tc.k)
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
