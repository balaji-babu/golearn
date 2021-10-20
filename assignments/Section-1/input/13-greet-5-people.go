package greet5People

import (
	"fmt"
	"os"
)

func Greet5People() {
	fmt.Printf("There are %d people! \n", len(os.Args)-1)
	count := 0
	for x := range os.Args {
		if count != 0 {
			fmt.Printf("Hello great %s \n", os.Args[x])
		}
		count++
	}
	fmt.Println("Nice to meet you all")
}
