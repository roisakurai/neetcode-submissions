func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	seen := make(map[rune]int)
	seen2 := make(map[rune]int)

	for _, ch := range s {
		seen[ch]++
	}

	for _, ch2 := range t {
		seen2[ch2]++
	}

	fmt.Println(seen)

	for _, a:= range t {
		if seen[a] != seen2[a] {
			return false
		}
	}

	return true
}
