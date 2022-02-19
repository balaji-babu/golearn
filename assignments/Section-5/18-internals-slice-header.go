package section5

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func ish() {
	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	array1 := [size]int{}

	// 2. print the memory usage (use the report func).
	report("Array1")

	// 3. copy the array to a new array.
	array2 := array1

	// 4. print the memory usage
	report("Array2")

	// 5. pass the array to the passArray function
	passArray(array1)

	// 6. convert one of the arrays to a slice
	slice1 := append([]int(nil), array2[:]...)
	// 7. slice only the first 1000 elements of the array
	slice2 := array1[:1000]

	// 8. slice only the elements of the array between 1000 and 10000
	slice3 := array1[1000:10000]
	// 9. print the memory usage (report func)
	report("Slice3 memory usage")
	// 10. pass the one of the slices to the passSlice function
	passSlice(slice1)
	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Printf("Size of Array1 %d\n", unsafe.Sizeof(array1))
	fmt.Printf("Size of Array2 %d\n", unsafe.Sizeof(array2))
	fmt.Printf("Size of Slice1 %d\n", unsafe.Sizeof(slice1))
	fmt.Printf("Size of Slice2 %d\n", unsafe.Sizeof(slice2))
	fmt.Printf("Size of Slice2 %d\n", unsafe.Sizeof(slice3))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
