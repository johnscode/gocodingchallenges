package allpalindromes

func expand(i int, j int, str string, palindromes map[string]string) map[string]string {
	runes := []rune(str)
	for i >= 0 && j < len(str) && runes[i] == runes[j] {
		pal := string(runes[i : j+1])
		palindromes[pal] = pal
		i -= 1
		j += 1
	}
	return palindromes
}

func FindAllPalindromes(str string) []string {
	palindromes := map[string]string{}

	for i := 0; i < len(str); i++ {
		palindromes = expand(i, i, str, palindromes)
	}
	for i := 0; i < len(str)-1; i++ {
		palindromes = expand(i, i+1, str, palindromes)
	}
	//fmt.Printf("palindromes: %+v", palindromes)
	res := []string{}
	for _, v := range palindromes {
		res = append(res, v)
	}
	return res
}
