package longestSubstringNoRepeat

func lengthOfLongestSubstring(s string) int {
	var running int
	seen := map[uint8]int{}
	length := 0

	if len(s) <= 1 {
		return len(s)
	}

	for pos := 0; pos < len(s); pos++ {
		val, ok := seen[s[pos]]
		if !ok {
			running++
		} else {
			pos = val + 1
			seen = map[uint8]int{}
			running = 1
		}
		if running > length {
			length = running
		}
		seen[s[pos]] = pos
	}
	if running > length {
		length = running
	}

	return length
}
