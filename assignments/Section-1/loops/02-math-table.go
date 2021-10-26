package loops

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	usageMsg = "Usage:[op=" + opt + "][size]"
	sizeMsg  = "size is missing"
	opt      = "*/+-%"
)

func mathTables() {
	switch l := len(os.Args); {
	case l == 2:
		fmt.Println(sizeMsg)
		fallthrough
	case l < 2:
		fmt.Println(usageMsg)
		return
	}

	if strings.IndexAny(os.Args[1], opt) == -1 {
		fmt.Println("Invalid Operator")
		fmt.Printf("Valid ops are one of: %s ", opt)
		return
	}

	size, err := strconv.Atoi(os.Args[2])

	if err != nil || size <= 0 {
		fmt.Println("Enter a valid Number")
		return
	}

	operatorFunction(os.Args[1], size)
}

func operatorFunction(op string, size int) {
	fmt.Printf("%5s", op)

	for i := 1; i < size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println("\n")
	for i := 1; i < size; i++ {
		fmt.Printf("%5d", i)
		for j := 1; j < size; j++ {
			var res int
			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}

			fmt.Printf("%5d", res)
		}
		fmt.Println("\n")
	}
}
