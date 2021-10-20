package printName

import (
	"fmt"
	"os"
)

func PrintName() {
	name := os.Args[1]
	fmt.Printf("Hi %s", name)
	fmt.Printf("How are you?")
}
