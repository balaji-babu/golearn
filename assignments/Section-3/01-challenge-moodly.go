package section3

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func challengeMoodly() {
	if len(os.Args) < 2 {
		fmt.Printf("Enter Your Name\n")
		fmt.Printf("eg: go run main.go Balaji\n")
		return
	}

	mood := [...]string{"feels good 👍", "feels bad 👎", "feels sad 😞", "feels happy 😀", "feels awesome 😎", "feels terrible 😩"}
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(mood))
	fmt.Printf("%s %s", os.Args[1], mood[n])

}
