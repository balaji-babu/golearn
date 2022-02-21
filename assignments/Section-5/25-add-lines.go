package section5

import (
	"fmt"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func addLines() {
	lyric := strings.Fields(`yesterday all my troubles seemed so far away now it looks as though they are here to stay oh I believe in yesterday`)

	fix := make([]string, len(lyric)+3)

	cutpoints := []int{8, 10, 5}

	for i, n := 0, 0; n < len(lyric); i++ {

		n += copy(fix[n+i:], lyric[n:n+cutpoints[i]])

		fix[n+i] = "\n"

	}

	s.Show("fix slice", fix)

	for _, w := range fix {
		fmt.Print(w)
		if w != "\n" {
			fmt.Print(" ")
		}
	}
}

func init() {
	//
	// YOU DON'T NEED TO TOUCH THIS
	//
	// This initializes some options for the prettyslice package.
	// You can change the options if you want.
	//
	// This code runs before the main function above.
	//
	// s.Colors(false)     // if your editor is light background color then enable this
	s.PrintBacking = true  // prints the backing arrays
	s.MaxPerLine = 5       // prints max 15 elements per line
	s.SpaceCharacter = '⏎' // print this instead of printing a newline (for debugging)
}
