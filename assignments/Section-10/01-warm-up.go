package section10

import "fmt"

var (
	phones      map[string]string
	products    map[int]bool
	multiPhones map[string][]string
	basket      map[int]map[int]int
)

func warmUp() {
	fmt.Printf("phones     : %#v\n", phones)
	fmt.Printf("products   : %#v\n", products)
	fmt.Printf("multiPhones: %#v\n", multiPhones)
	fmt.Printf("basket     : %#v\n", basket)
}
