package printf

import (
	"fmt"
	"os"
)

func printName() {
	msg := "My First name is %s, and my last name is %s"

	fmt.Printf(msg, os.Args[1], os.Args[2])
}
