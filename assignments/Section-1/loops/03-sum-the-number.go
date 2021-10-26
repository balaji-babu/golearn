package loops

import "fmt"

func sumTheNumber() {
	y := 0
	for x := range [11]int{} {
		y = y + x
	}
	fmt.Println(y)
}
