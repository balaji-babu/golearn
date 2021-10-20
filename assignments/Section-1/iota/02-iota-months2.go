package iota

import "fmt"

const (
	Jan = iota + 1
	Feb
	Mar
	Apr
	May
	June
	July
	Aug
	Sep
	Oct
	Nov
	Dec
)

func IotaMonths1() {
	fmt.Println(Jan, Feb, Mar, Apr, May, June, July, Aug, Sep, Oct, Nov, Dec)
}
