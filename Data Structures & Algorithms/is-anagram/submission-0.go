func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	word1, word2 := make(map[rune]int), make(map[rune]int)
	for i, ch := range s {
		word1[ch]++
		word2[rune(t[i])]++
	}

	for k, v := range word1 {
		if word2[k] != v {
			return false
		}
	}
	return true


}
