package section3

import (
	"fmt"
	"os"
	"strings"
)

const errMsg = "Tell me a book title"

var status bool

func hipstersLoveSearch() {
	bt := [...]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	if len(os.Args) != 2 {
		fmt.Printf("%q\n", errMsg)
		return
	}

	for _, v := range bt {
		if strings.Contains(strings.ToLower(v), strings.ToLower(os.Args[1])) {
			fmt.Println("+", v)
			status = true
		}
	}
	if !status {
		fmt.Printf("We don't have the book: %q\n", os.Args[1])
	}

}
