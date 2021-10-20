package iota

import "fmt"

const (
	Spring = 3 * (iota + 1)
	Summer
	Fall
	Winter
)

func IotaSeasons() {
	fmt.Println(Winter, Spring, Summer, Fall)
}
