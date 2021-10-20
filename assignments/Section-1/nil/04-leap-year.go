package nil

import (
	"fmt"
	"os"
	"strconv"
)

func leapYear() {
	if len(os.Args) < 2 {
		fmt.Println("Enter year")
	} else if year, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("Enter a valid year.")
	} else if year%400 == 0 || year%4 == 0 {
		fmt.Printf("%d is a leap year", year)
	} else {
		fmt.Printf("%d is not a leap year", year)
	}

}
