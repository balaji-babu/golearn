package ifElse

import "fmt"

func AgeSeasons() {
	var age int = 10
	if age > 60 {
		fmt.Println("Getting Older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young Blood")
	} else {
		fmt.Println("Booting up")
	}
}
