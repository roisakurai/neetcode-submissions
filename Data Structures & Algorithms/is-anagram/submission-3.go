func isAnagram(s string, t string) bool {
	words1 := make(map[rune]int)
	words2 := make(map[rune]int)

	if len(s) != len(t) {
		return false
	}

	for _, a := range s {
		words1[a]++
	}

	for _, b := range t {
		words2[b]++
	}

	for _, c := range s {
		if words1[c] != words2[c] {
			return false
		}
	}


	return true
}
