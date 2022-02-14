package section5

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	intData []int
	even    []int
	odd     []int
)

func sliceByNumbers() {
	data := "2 4 6 1 3 5"

	// convert string to []int
	for _, value := range strings.Fields(data) {
		value, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("Error 😥😥😥")
			fmt.Printf("%s", err)
			return
		}
		intData = append(intData, value)
	}
	fmt.Printf("nums:       	 %v\n", intData)

	even, odd = intData[:3], intData[3:]

	fmt.Printf("evens:      	 %v\n", even)
	fmt.Printf("odds:       	 %v\n", odd)
	fmt.Printf("middle:     	 %v\n", intData[2:4])
	fmt.Printf("first 2:    	 %v\n", intData[0:2])
	fmt.Printf("last 2:       	 %v\n", intData[len(intData)-2:])
	fmt.Printf("even last 1:     %v\n", even[2:])
	fmt.Printf("odd last 2:      %v\n", odd[1:])
}
