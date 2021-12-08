package section3

import (
	"fmt"
	"os"
	"strconv"
)

const (
	errMsg = "Please tell me numbers (maximum 5 numbers)."
)

func average() {

	args := os.Args[1:]
	var count, sum float64

	if l := len(args); l < 1 || l > 5 {
		fmt.Println(errMsg)
		return
	}

	for i, v := range args {
		value, err := strconv.ParseFloat(v, 64)
		if err != nil {
			args[i] = "0"
			continue
		}
		count++
		sum += value
	}
	fmt.Printf("Your Numbers:%s\n", args)
	fmt.Printf("Average: %.2f\n", sum/count)
}
