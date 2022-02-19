package section5

import (
	"fmt"
	"sort"
)

func ibas() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	sortData := items[(len(items)/2)-1 : (len(items)/2)+2]

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE\
	sort.Sort(sort.StringSlice(sortData))
	fmt.Println()
	fmt.Println("Sorted  :", items)
}
