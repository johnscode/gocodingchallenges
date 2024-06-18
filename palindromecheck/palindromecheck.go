package palindromecheck

import "strings"

func PalindromeCheck(str string) bool {
	trimmedStr := strings.ReplaceAll(str, " ", "")
	trimmedStrLen := len(trimmedStr)
	chars := []rune(trimmedStr)
	for i := 0; i < trimmedStrLen/2; i++ {
		if chars[i] != chars[trimmedStrLen-i-1] {
			return false
		}
	}
	return true
}
