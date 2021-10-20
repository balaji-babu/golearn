package conversion

import "fmt"

func convert1() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}
