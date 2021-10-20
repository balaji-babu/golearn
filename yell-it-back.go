package yellItBack

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func YellItBack() {
	fmt.Println("Please Enter Your Name")

	name := os.Args[1]

	output := strings.ToUpper(name) + strings.Repeat("!", utf8.RuneCountInString(name))

	fmt.Printf("Hello %s \n", output)
}
