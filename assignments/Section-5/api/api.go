package api

import (
	"fmt"
	"math/rand"
	"runtime"
)

var temps = []int{5, 10, 3, 25, 45, 80, 90}

// Read returns a slice of elements from the temps slice.
func Read(start, stop int) []int {
	// ----------------------------------------
	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA

	portion := temps[start:stop:stop]

	// RESTRICTIONS — ONLY ADD YOUR CODE IN THIS AREA
	// ----------------------------------------

	return portion
}

// All returns the temps slice
func All() []int {
	return temps
}

func Read1() []int {
	return rand.Perm(2 << 22)
}

func Report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
