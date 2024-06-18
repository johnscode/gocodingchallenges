package main

import (
	"fmt"
	"gocoding/uniquestring"
)

func main() {
	str := "Hello, World!"
	maxUnique := uniquestring.CountUniqueChars(str)

	fmt.Printf("The string \"%s\" contains %d unique characters.\n", str, maxUnique)
}
