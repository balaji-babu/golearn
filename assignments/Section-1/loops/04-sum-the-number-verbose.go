package loops

import (
	"fmt"
	"strconv"
)

const (
	max = 11
)

func sumTheNumberVerbose() {
	y := 0
	for x := range [max]int{} {
		y = y + x
		if x != 0 {
			fmt.Print(x)
			if x < max-1 {
				fmt.Print(" + ")
			}

		}

	}
	fmt.Printf(" = %s", strconv.Itoa(y))
}
