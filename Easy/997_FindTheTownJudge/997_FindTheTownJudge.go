package findTheTownJudge

// in a town, there are n people labeled from 1 to n.
// There is a rumor that one of these people is secretly the town judge.
//
// If the town judge exists, then:
//    The town judge trusts nobody.
//    Everybody (except for the town judge) trusts the town judge.
//    There is exactly one person that satisfies properties 1 and 2.
//
// You are given an array trust where trust[i] = [ai, bi]
// representing that the person labeled ai trusts the person labeled bi.
//
// Return the label of the town judge if the town judge exists and can be identified,
// or return -1 otherwise.

func findJudge(n int, trust [][]int) int {
	candidates := make([]int, 0)
	invalids := map[int]struct{}{}
	trustNums := map[int]int{}

	if n == 1 {
		return 1
	}

	for _, val := range trust {
		if _, ok := trustNums[val[1]]; !ok {
			trustNums[val[1]] = 1
		} else {
			trustNums[val[1]]++
		}
		invalids[val[0]] = struct{}{}
	}

	for key, val := range trustNums {
		if val == (n - 1) {
			if _, ok := invalids[key]; !ok {
				candidates = append(candidates, key)
			}
		}
	}

	if len(candidates) != 1 {
		return -1
	}
	return candidates[0]
}
