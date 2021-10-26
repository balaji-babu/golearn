package loops

import (
	"fmt"
	"os"
	"strconv"
)

func dynamicTable() {
	if len(os.Args) != 2 {
		fmt.Println("Provide valid size of the table")
		return
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil || size <= 0 {
		fmt.Printf("%s", "provide valid integer")
		return
	}
	fmt.Printf("%5s", "*")

	for i := 1; i < size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println("\n")
	for i := 1; i < size; i++ {
		fmt.Printf("%5d", i)
		for j := 1; j < size; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println("\n")
	}

}
