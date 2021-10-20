package nil

import (
	"fmt"
	"os"
	"strconv"
)

func feetToMeter() {
	feet, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a number \n", os.Args[1])
		return
	}
	meter := float64(feet) / 3.2808

	fmt.Printf("%s feet is %.2f meters.", os.Args[1], meter)

}
