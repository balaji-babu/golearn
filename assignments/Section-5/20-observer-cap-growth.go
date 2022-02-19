package section5

import (
	"fmt"
	"strconv"
)

func ocg() {
	slice1 := []string(nil)
	var oldCap float64

	for i := range [1e7]int{} {
		c := float64(cap(slice1))
		slice1 = append(slice1, strconv.Itoa(i))
		if c == 0 || c != oldCap {
			fmt.Printf("len:%d,\tcap:%1.2f,\tgrowth:%1.2f\n", len(slice1), c, c/oldCap)
		}
		oldCap = c
	}
}
