package section3

import (
	"fmt"
	"strings"
)

func wizardPrinter() {
	wp := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for ir := range wp {
		if ir == 0 {
			fmt.Println(strings.Repeat("=", 50) + "\n")
		}
		fmt.Printf("%-15s %-15s %-15s\n", wp[ir][0], wp[ir][1], wp[ir][2])
	}
}
