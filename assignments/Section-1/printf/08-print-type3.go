package printf

import "fmt"

func PrintType3() {
	fmt.Printf("Type of %s is %[1]T\n", "hello")
}
