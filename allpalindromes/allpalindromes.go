package allpalindromes

import (
	"gocoding/palindromecheck"
	"gocoding/uniquecombos"
)

func FindAllPalindromes(str string) []string {
	allPalindromes := []string{}
	uniqueStrings := uniquecombos.FindUniqueCombinations(str)
	for _, uniqueString := range uniqueStrings {
		if palindromecheck.PalindromeCheck(uniqueString) {
			allPalindromes = append(allPalindromes, uniqueString)
		}
	}
	return allPalindromes
}
