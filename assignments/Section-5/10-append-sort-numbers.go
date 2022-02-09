package section5

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

var input []int

func appendSortNum() {
	if len(os.Args) <= 1 {
		fmt.Println("Provide a few numbers.")
		return
	}

	for _, v := range os.Args[1:] {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		input = append(input, n)
	}
	sort.Ints(input)
	fmt.Printf("%d", input)
}
