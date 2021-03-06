package string

func bfSearch(main, pattern string) int {
	if len(main) == 0 || len(pattern) == 0 || len(main) < len(pattern) {
		return -1
	}

	for i := 0; i < len(main) - len(pattern); i++ {
		subStr := main[i: i+len(pattern)]
		if subStr == pattern {
			return i
		}
	}

	return -1
}
