package section5

import (
	"fmt"
	"golearn/assignments/Section-5/api"
)

func ltbas() {
	// DO NOT CHANGE ANYTHING IN THIS CODE.

	// get the first three elements from api.temps
	received := api.Read(0, 3)

	// append changes the api package's temps slice's
	// backing array as well.
	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)
}
