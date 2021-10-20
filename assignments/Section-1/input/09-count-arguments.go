package countArg

import (
	"fmt"
	"os"
)

func CountArgs() {
	count := 0

	count = len(os.Args) - 1

	fmt.Printf("There are %d names.\n", count)
}
