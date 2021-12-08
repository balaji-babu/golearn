package section3

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func challengeMoodly() {
	if len(os.Args) != 3 {
		fmt.Printf("Enter Your Name and Mood\n")
		fmt.Printf("eg: go run main.go Balaji Positive/Negative\n")
		return
	}
	if os.Args[2] != strings.ToLower("positive") && os.Args[2] != strings.ToLower("negative") {
		fmt.Printf("Please Enter Positive/Negative value for second Args\n")
		fmt.Printf("eg: go run main.go Balaji Positive/Negative\n")
		return
	}

	moods := [2][3]string{
		{"feels good 👍", "feels happy 😀", "feels awesome 😎"},
		{"feels bad 👎", "feels sad 😞", "feels terrible 😩"},
	}

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(moods[0]))

	var mood int
	if strings.ToLower(os.Args[2]) != "positive" {
		mood = 1
	}
	fmt.Printf("%s %s\n", os.Args[1], moods[mood][n])
}
