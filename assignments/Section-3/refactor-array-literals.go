package section3

import (
	"fmt"
	"strings"
)

func refactorArrayLiterals() {
	names := [3]string{"A", "B", "C"}
	distances := [5]int{50, 100, 150, 200, 250}
	data := [5]byte{'C', 'O', 'A', 'S', 'T'}
	ratios := [1]float64{3.14}
	alives := [4]bool{true, false, true, false}
	var zero [0]byte

	fmt.Println(strings.Repeat("+", 10) + " Names: " + strings.Repeat("+", 10))
	for x := 0; x < len(names); x++ {
		fmt.Printf("names[%d]: %q\n", x, names[x])
	}
	fmt.Println(strings.Repeat("+", 10) + " Distances: " + strings.Repeat("+", 10))
	for x := 0; x < len(distances); x++ {
		fmt.Printf("distances[%d]: %d\n", x, distances[x])
	}
	fmt.Println(strings.Repeat("+", 10) + " Data: " + strings.Repeat("+", 10))
	for x := 0; x < len(data); x++ {
		fmt.Printf("data[%d]: %d\n", x, data[x])
	}
	fmt.Println(strings.Repeat("+", 10) + " Ratios: " + strings.Repeat("+", 10))
	for x := 0; x < len(ratios); x++ {
		fmt.Printf("ratios[%d]: %.2f\n", x, ratios[x])
	}
	fmt.Println(strings.Repeat("+", 10) + " Alives: " + strings.Repeat("+", 10))
	for x := 0; x < len(alives); x++ {
		fmt.Printf("alives[%d]: %t\n", x, alives[x])
	}
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println(strings.Repeat("+", 10) + " Using For Range: " + strings.Repeat("+", 10))
	fmt.Println(strings.Repeat("=", 20))

	fmt.Println(strings.Repeat("+", 10) + " Names: " + strings.Repeat("+", 10))
	for k, v := range names {
		fmt.Printf("names[%d]: %q\n", k, v)
	}
	fmt.Println(strings.Repeat("+", 10) + " Distances: " + strings.Repeat("+", 10))
	for k, v := range names {
		fmt.Printf("distances[%d]: %d\n", k, v)
	}
	fmt.Println(strings.Repeat("+", 10) + " Data: " + strings.Repeat("+", 10))
	for k, v := range names {
		fmt.Printf("data[%d]: %d\n", k, v)
	}
	fmt.Println(strings.Repeat("+", 10) + " Ratios: " + strings.Repeat("+", 10))
	for k, v := range names {
		fmt.Printf("ratios[%d]: %.2f\n", k, v)
	}
	fmt.Println(strings.Repeat("+", 10) + " Alives: " + strings.Repeat("+", 10))
	for k, v := range names {
		fmt.Printf("alives[%d]: %t\n", k, v)
	}
	fmt.Println(strings.Repeat("+", 10) + " zero: " + strings.Repeat("+", 10))
	for k, v := range zero {
		fmt.Printf("zero[%d]: %t\n", k, v)
	}
}
