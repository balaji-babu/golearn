package section3

import (
	"fmt"
	"os"
	"strconv"
)

const (
	errMsg = "Please tell me numbers (maximum 5 numbers)."
)

func sorter() {
	args := os.Args[1:]

	var na [5]int

	switch l := len(args); {
	case l == 0:
		fmt.Println("Please give me up to 5 numbers.")
		return
	case l > 5:
		fmt.Println("Sorry. Go arrays are fixed.",
			"So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	for i, v := range args {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			continue
		}

		na[i] = int(value)
	}
	for range na {
		for i := range na {
			if i < len(na)-1 && na[i] > na[i+1] {
				na[i], na[i+1] = na[i+1], na[i]
			}
		}

	}
	fmt.Printf("%d", na)
}
