func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	output := ""

	for i := range lower {
		if (lower[i] >= 'a' && lower[i] <= 'z') || (lower[i] >= '0' && lower[i] <= '9') {
			output += string(lower[i])
		}
	}
	
	for i := 0; i < len(output)/2;i++ {
		if output[i] != output[len(output)-1-i] {
			return false
		}
	}

	return true
}
