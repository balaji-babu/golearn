package section3

import (
	"fmt"
	"strings"
)

func assign() {
	var books [4]string = [4]string{"Common Sense Investing", "Beating The Street", "Common Sense on Mutual Funds", "The Intelligent Investor"}

	booksUpper, bookslower := books, books

	for k, v := range books {
		booksUpper[k] = strings.ToUpper(v)
		bookslower[k] = strings.ToLower(v)
	}

	fmt.Println("Books")
	fmt.Printf("Books: %q\n", books)
	fmt.Println("Books UpperCase")
	fmt.Printf("Upper: %q\n", booksUpper)
	fmt.Println("Books LowerCase")
	fmt.Printf("Lower: %q\n", bookslower)
}
