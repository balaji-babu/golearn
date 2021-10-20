package printf

import "fmt"

func PrintType1() {
	fmt.Printf("Type of %[1]d is %[1]T\n", 3)
}
