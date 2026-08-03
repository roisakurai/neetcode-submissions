func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	output := ""
	
	for i := range lower {
		if (lower[i] >= 'a' && lower[i] <= 'z') || (lower[i] >= '0' && lower[i] <= '9') {
			output += string(lower[i])
		}
	}

	half := len(output)/2
	
	for i := range half {
		if output[i] != output[len(output)-1-i] {
			return false
		}
	}

	return true
}
