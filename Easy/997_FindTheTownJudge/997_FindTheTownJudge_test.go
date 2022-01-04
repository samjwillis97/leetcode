package findTheTownJudge

import (
	"fmt"
	"testing"
)

func TestFindTheTownJudgeExamples(t *testing.T) {
	tests := []struct {
		n     int
		trust [][]int
		want  int
	}{
		{2, [][]int{{1, 2}}, 2},
		{3, [][]int{{1, 3}, {2, 3}}, 3},
		{3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("findJudge=%d", i), func(t *testing.T) {
			got := findJudge(tc.n, tc.trust)
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func TestFindTheTownJudgeCases(t *testing.T) {
	tests := []struct {
		n     int
		trust [][]int
		want  int
	}{
		{1, [][]int{{}}, 1},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("findJudge=%d", i), func(t *testing.T) {
			got := findJudge(tc.n, tc.trust)
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
