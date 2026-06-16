func isAlphaNum(c byte) bool {
	return ('a'<=c && c<='z') || ('A'<=c && c<='Z') || ('0'<=c && c<='9')
}

func isPalindrome(s string) bool {
	right := 0
	left := len(s)-1
	for left > right {
		if !isAlphaNum(s[left]) {
			left--
			continue
		}
		if !isAlphaNum(s[right]) {
			right++
			continue
		}
		lowerCharLeft := strings.ToLower(string(s[left]))[0]
		lowerCharRight := strings.ToLower(string(s[right]))[0]
		if lowerCharLeft!= lowerCharRight {
			return false
		}
		left--
		right++
	}
	return true
}