package section6

import (
	"fmt"
	"os"
)

func sortToAFile2() {
	var output []byte
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Send me some items and I will sort them")
		return
	}

	for _, v := range args {
		output = append(output, v...)
		output = append(output, '\n')
	}

	WriteToFile(output)
}
