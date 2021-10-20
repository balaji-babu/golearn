package countArg

import (
	"fmt"
	"os"
)

func PrintPath() {
	fmt.Println(os.Args[0])
}
