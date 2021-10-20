package countTheCharacters

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func countTheCharacters() {
	name := os.Args[1]
	fmt.Println(utf8.RuneCountInString(name))
}
