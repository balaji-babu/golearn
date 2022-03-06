package section8

import "fmt"

func convert() {
	words := []string{
		"gopher",
		"programmer",
		"go language",
		"go standard library",
	}

	var bwords [][]byte
	//   1. Loop over the words slice
	for _, v := range words {
		strConv := []byte(v)
		fmt.Printf("%d\n", strConv)
		bwords = append(bwords, strConv)
	}

	for _, v := range bwords {
		fmt.Printf("%s\n", v)
	}
}
