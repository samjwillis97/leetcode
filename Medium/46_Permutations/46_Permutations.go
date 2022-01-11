package permutations

func permute(nums []int) [][]int {
	result := make([][]int, 0)

	var perm func(start int, curPerm []int)

	perm = func(start int, curPerm []int) {
		if start == len(nums)-1 {
			dst := make([]int, len(nums))
			copy(dst, curPerm)
			result = append(result, dst)
		}

		for i := start; i < len(nums); i++ {
			curPerm[start], curPerm[i] = curPerm[i], curPerm[start]
			perm(start+1, curPerm)
			curPerm[start], curPerm[i] = curPerm[i], curPerm[start]
		}
		return
	}
	perm(0, nums)
	return result
}
