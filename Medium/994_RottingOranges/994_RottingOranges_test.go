package rottingOranges

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRottingOranges(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("UpdateMatrix=%d", i), func(t *testing.T) {
			printMatrix(tc.matrix)
			got := orangesRotting(tc.matrix)
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
