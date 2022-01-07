package floodFill

import (
	"fmt"
	"testing"
)

// image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, newColor = 2
//Output: [[2,2,2],[2,2,0],[2,0,1]]

func TestFloodFill(t *testing.T) {
	tests := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		want     [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2, [][]int{{2, 2, 2}, {2, 2, 2}}},
		{[][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1, [][]int{{0, 0, 0}, {0, 1, 1}}},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("FloodFill=%d", i), func(t *testing.T) {
			got := floodFill(tc.image, tc.sr, tc.sc, tc.newColor)
			if !testEq(tc.want, got) {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func testEq(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
