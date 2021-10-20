package nil

import (
	"fmt"
	"os"
	"strconv"
)

func oddOrEven() {
	divisibleValue := 8

	if len(os.Args) < 2 {
		fmt.Println("Enter a number")
	} else if n, err := strconv.Atoi(os.Args[1]); n < 0 || err != nil {
		fmt.Println("Enter a positivie integer")
	} else if n&1 != 0 {
		fmt.Printf("%d is an odd number\n", n)
	} else if n&1 == 0 && n%divisibleValue == 0 {
		fmt.Printf("%d is an even number and it's divisible by %d\n", n, divisibleValue)
	} else {
		fmt.Printf("%d is an even number\n", n)
	}

}
