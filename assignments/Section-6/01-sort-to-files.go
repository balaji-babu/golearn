package section6

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

func sortToAFile() {
	var output []byte
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Send me some items and I will sort them")
		return
	}
	sort.Strings(args)
	for _, v := range args {
		output = append(output, v...)
		output = append(output, '\n')
	}

	WriteToFile(output)
}

func WriteToFile(name []byte) {
	ioutil.WriteFile("./sorted.txt", name, 0644)
}
