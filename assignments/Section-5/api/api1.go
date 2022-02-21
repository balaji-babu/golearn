package api

import (
	"fmt"
	"math/rand"
	"runtime"
)

func Read() []int {
	return rand.Perm(2 << 22)
}

func Report() {
	var m runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m)
	fmt.Printf(" > Memory Usage: %v KB\n", m.Alloc/1024)
}
