package section5

import (
	"fmt"
	"golearn/assignments/Section-5/api"
	"io/ioutil"
)

func ftml() {
	api.Report()
	millions := api.Read1()
	last10 := make([]int, 10)
	copy(last10, millions[len(millions)-10:])
	millions = last10
	fmt.Printf("\nLast 10 elements: %d\n\n", last10)
	api.Report()
	fmt.Fprintln(ioutil.Discard, millions[0])
}
