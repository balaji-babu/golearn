package loops

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usageMsg string = "Enter Two Integers Min and Max Value"
)

func onlyEvens() {
	switch l := len(os.Args); {
	case l == 2:
		fmt.Println("Max value is missing")
		fallthrough
	case l < 2:
		fmt.Println(usageMsg)
		return
	}

	var (
		min, max, sum int
		err, err2     error
	)

	min, err = strconv.Atoi(os.Args[1])
	max, err2 = strconv.Atoi(os.Args[2])

	if err != nil || err2 != nil || min >= max {
		fmt.Println("Please make sure, Min and Max value is Integer and Min is less than Max")
		return
	}
	for x := min; x <= max; x++ {
		if x%2 != 0 {
			continue
		}
		sum = sum + x
		fmt.Print(x)

		if x < max-1 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d", sum)
}
