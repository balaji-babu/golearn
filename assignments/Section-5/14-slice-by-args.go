package section5

import (
	"fmt"
	"os"
	"strconv"
)

var (
	argsValue  []string
	start, end int
	err        error
)

const (
	printMsg = "Provide only the [Starting] and [Stopping] Position"
	errMsg   = "Error Message: We expect only two arguments"
	intMsg   = "Please insert only positive integer values"
)

func sliceByArgs() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	if len(os.Args) == 1 {
		fmt.Printf("%q\n", ships)
		fmt.Printf("%q\n", printMsg)
	}

	start, end = 0, len(ships)

	switch value := os.Args[1:]; len(value) {
	default:
		fallthrough
	case 0:
		fmt.Printf("%q\n", printMsg)
		return
	case 2:
		end, err = strconv.Atoi(value[1])
		if err != nil {
			fmt.Printf("%q\n", intMsg)
			fmt.Printf("%q\n", printMsg)
			return
		}
		fallthrough
	case 1:
		start, err = strconv.Atoi(value[0])
		if err != nil {
			fmt.Printf("%q\n", intMsg)
			fmt.Printf("%q\n", printMsg)
			return
		}
	}

	if l := len(ships); start < 0 || start > l || end > l || start > end {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Printf("%q", ships[start:end])
}
