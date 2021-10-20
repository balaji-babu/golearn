package challenge

import (
	"fmt"
	"os"
)

func challenge01() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: [username] [password]")
	} else if os.Args[1] == "jack" && os.Args[2] == "1888" {
		fmt.Printf("Access granted for %q.", os.Args[1])
	} else if os.Args[1] != "jack" {
		fmt.Printf("Access denied for %q.", os.Args[1])
	} else {
		fmt.Printf("Invalid password for %q.", os.Args[1])
	}
}
