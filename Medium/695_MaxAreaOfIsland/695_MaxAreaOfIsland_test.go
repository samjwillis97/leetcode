package maxAreaOfIsland

import (
	"fmt"
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}, 6},
		{[][]int{{0, 0, 0, 0, 0, 0, 0, 0}}, 0},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("MaxAreaOfIsland=%d", i), func(t *testing.T) {
			got := maxAreaOfIsland(tc.grid)
			if got != tc.want {
				t.Fatalf("got  %v;\nwant %v;", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
