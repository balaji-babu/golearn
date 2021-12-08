package section3

import (
	"fmt"
	"strings"
)

var (
	names     [3]string
	distances [5]int
	data      [5]byte
	ratios    [1]float64
	alives    [4]bool
	zero      [0]byte
)

func getSetArrayElements() {
	// Assing Best Friends Name to Names Array

	names[0] = "A"
	names[1] = "B"
	names[2] = "C"

	// Assign distances to the closest cities to you to the distance array

	distances[0] = 100
	distances[1] = 200
	distances[2] = 300
	distances[3] = 400
	distances[4] = 500

	// Assign arbitary bytes to data array
	data[0] = 'C'
	data[1] = 'O'
	data[2] = 'M'
	data[3] = 'I'
	data[4] = 'C'

	//Assign a value to ratios array
	ratios[0] = 3.14

	// Assign True or False Values to alives array
	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

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
}
