package greetMorePeople

import (
	"fmt"
	"os"
)

func GreetMorePeople() {
	noOfArgs := len(os.Args) - 1
	fmt.Printf("There are %d people!\n", noOfArgs)
	count := 0
	for x := range os.Args {
		if count != 0 {
			fmt.Printf("Hello great %s \n", os.Args[x])
		}
		count++
	}

}
