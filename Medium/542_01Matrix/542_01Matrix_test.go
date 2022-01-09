package matrix01

import (
	"fmt"
	"reflect"
	"testing"
)

func Test01Matrix(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
		{[][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 1, 1, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0, 1, 0}, {1, 1, 1, 1, 0, 1, 0, 0, 1, 1}}, [][]int{{1, 0, 1, 1, 0, 0, 1, 0, 0, 1}, {0, 1, 1, 0, 1, 0, 1, 0, 1, 1}, {0, 0, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 0, 0, 0, 1}, {0, 0, 1, 0, 1, 1, 1, 0, 1, 0}, {0, 1, 0, 1, 0, 1, 0, 0, 1, 1}, {1, 0, 0, 0, 1, 2, 1, 1, 0, 1}, {2, 1, 1, 1, 1, 2, 1, 0, 1, 0}, {3, 2, 2, 1, 0, 1, 0, 0, 1, 1}}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("UpdateMatrix=%d", i), func(t *testing.T) {
			printMatrix(tc.matrix)
			got := updateMatrix(tc.matrix)
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
