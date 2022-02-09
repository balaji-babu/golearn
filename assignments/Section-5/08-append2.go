package section5

import (
	"fmt"
	"time"
)

var (
	pizza     []string
	departure []time.Time
	student   []int
	lights    []bool
)

func append2() {
	pizza = append(pizza, "pepperoni", "onions", "extra", "cheese")
	departure = append(departure, time.Now(), time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*48), time.Now().Add(time.Hour*72))
	student = append(student, 1991, 1994, 2000, 2012, 2016)
	lights = append(lights, true, false, true, false)

	fmt.Printf("pizza       : %s\n", pizza)
	fmt.Printf("departure   : %s\n", departure)
	fmt.Printf("student     : %d\n", student)
	fmt.Printf("lights      : %t\n", lights)
}
