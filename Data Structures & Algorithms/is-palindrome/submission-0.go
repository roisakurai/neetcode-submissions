func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	result := ""

	for i := 0; i < len(lower); i++ {
		if ('a' <= lower[i] && lower[i] <= 'z') || ('0' <= lower[i] && lower[i] <= '9') {
			result = result + string(lower[i])
		}

	}

	for i := 0; i < len(result)/2; i++ {
		if result[i] != result[len(result)-1-i] {
			return false
		}
	}

	return true
}