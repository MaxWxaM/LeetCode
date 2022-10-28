func longestPalindrome(s string) string {

	if len(s) == 0 {
		return ""
	}
	maxSize := 1
	finalLeft := 0
	for i := 0; i <= len(s); i++ {
		oddLeft := i
		evenLeft := i
		if i+2 <= len(s)-1 && s[i] == s[i+2] {
			right := i + 3
			tempMaxSize := 3
			if tempMaxSize > maxSize {
				maxSize = tempMaxSize
				finalLeft = evenLeft
			}
			evenLeft -= 1
			for evenLeft >= 0 && right < len(s) && s[evenLeft] == s[right] {
				tempMaxSize += 2
				if tempMaxSize > maxSize {
					maxSize = tempMaxSize
					finalLeft = evenLeft
				}
				evenLeft -= 1
				right += 1
			}
		}
		if i+1 <= len(s)-1 && s[i] == s[i+1] {
			right := i + 2
			tempMaxSize := 2
			if tempMaxSize > maxSize {
				maxSize = tempMaxSize
				finalLeft = oddLeft
			}
			oddLeft -= 1
			for oddLeft >= 0 && right < len(s) && s[oddLeft] == s[right] {
				tempMaxSize += 2
				if tempMaxSize > maxSize {
					maxSize = tempMaxSize
					finalLeft = oddLeft
				}
				oddLeft -= 1
				right += 1
			}
		}
	}
	if maxSize == 1 {
		return string(s[finalLeft])
	} else {
		return string(s[finalLeft : finalLeft+maxSize])
	}
}