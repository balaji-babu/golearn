package iota

import "fmt"

const (
	Nov = 11 - iota
	Oct
	Sep
)

func IotaMonths1() {
	fmt.Println(Nov, Oct, Sep)
}
