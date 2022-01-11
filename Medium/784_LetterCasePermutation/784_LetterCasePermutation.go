package letterCasePermutation

import (
	"strconv"
	"strings"
)

func letterCasePermutation(s string) []string {
	result := make([]string, 0)

	var perm func(start int, curPerm string)

	//recurs := 0
	perm = func(index int, curPerm string) {
		//recurs++
		//current := recurs
		//fmt.Printf("R: %d, perm(%d, %s)\n", current, index, curPerm)
		if len(curPerm) == len(s) {
			//fmt.Printf("R: %d, Result: %s\n", current, curPerm)
			result = append(result, curPerm)
		} else {
			if isAlpha(s[index]) {
				perm(index+1, curPerm+swapCase(s[index]))
			}
			perm(index+1, curPerm+string(s[index]))
		}
		//fmt.Printf("R: %d, Returned\n", current)
		return
	}
	perm(0, "")
	return result
}

func isAlpha(u uint8) bool {
	if _, err := strconv.Atoi(string(u)); err != nil {
		return true
	}
	return false
}

func swapCase(u uint8) string {
	upd := strings.ToUpper(string(u))
	if string(u) == upd {
		return strings.ToLower(string(u))
	}
	return upd
}

func swapCaseElement(s string, i int) (string, bool) {
	if _, err := strconv.Atoi(string(s[i])); err == nil {
		return s, false
	}
	newStr := s[:i] + strings.ToUpper(string(s[i])) + s[i+1:]
	if newStr == s {
		return s[:i] + strings.ToLower(string(s[i])) + s[i+1:], true
	}
	return newStr, true
}
