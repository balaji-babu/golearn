package printf

import "fmt"

func PrintType2() {
	fmt.Printf("Type of %.2[1]f is %[1]T\n", 3.14)
}
