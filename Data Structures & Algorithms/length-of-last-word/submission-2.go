func lengthOfLastWord(s string) int {
	counts := 0

	if len(s) == 1 {
		return 1
	}

	for i := len(s)-1; i >= 0; i-- {
		if string(s[i]) != " " {
			counts++

			if string(s[i-1]) == " " {
				return counts	
			}
		}
	}

	return counts

}