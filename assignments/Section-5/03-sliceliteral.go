package section5

import "fmt"

var (
	names     []string
	distances []int
	data      []uint8
	ratios    []float64
	alives    []bool
)

func sliceLiteral() {

	names = []string{"AAA", "BBB", "CCC"}
	distances = []int{10, 20, 30, 40, 50}
	data = []uint8{'a', 'b', 'c', 'd', 'e'}
	ratios = []float64{3.14, 10.1}
	alives = []bool{true, false, false, true}

	fmt.Printf("names: %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives: %T %d %t\n", alives, len(alives), alives == nil)

	if len(distances) == len(data) {
		fmt.Printf("The length of the distances and the data slices are the same.")
	}
}
