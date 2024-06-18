package uniquestring

//func main() {
//	str := "Hello, World!"
//	maxUnique := maxUniqueCharacters(str)
//	fmt.Printf("The string \"%s\" contains %d unique characters.\n", str, maxUnique)
//}

func CountUniqueChars(str string) int {
	charSet := make(map[rune]bool)

	for _, char := range str {
		charSet[char] = true
	}

	return len(charSet)
}
