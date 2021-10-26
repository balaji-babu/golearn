package switchStatement

import (
	"fmt"
	"os"
	"strings"
)

const usage = `[command] [string]
Available commands: lower, upper and title`

func stringManipulator() {
	if len(os.Args) != 3 {
		fmt.Printf("%s", usage)
		return
	}
	switch os.Args[1] {
	case "lower":
		fmt.Printf(strings.ToLower(os.Args[2]))
	case "upper":
		fmt.Printf(strings.ToUpper(os.Args[2]))
	case "title":
		fmt.Printf(strings.Title(os.Args[2]))
	default:
		fmt.Printf("Unknown command: %q\n", os.Args[1])
	}
}
