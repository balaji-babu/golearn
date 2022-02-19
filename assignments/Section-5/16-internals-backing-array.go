package section5

import "fmt"

var mine []int

func iba() {
	nums := []int{56, 89, 15, 25, 30, 50}

	mine = append(mine, nums[:3]...)

	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine)
	fmt.Println("Original nums:", nums[:3])
}
