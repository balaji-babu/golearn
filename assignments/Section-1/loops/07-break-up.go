package loops

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usageMsg string = "Enter Two Integers Min and Max Value"
)

func breakUP() {
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
	x := min
	for {
		if x > max {
			break
		} else if x%2 != 0 {
			x++
			continue
		}

		fmt.Print(x)

		if x < max-1 {
			fmt.Print(" + ")
		}
		sum = sum + x
		x++
	}
	fmt.Printf(" = %d", sum)

}
