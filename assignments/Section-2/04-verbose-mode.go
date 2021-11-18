package luckyGuess

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxGuess = 5
	usage    = `
Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, harder it gets.

Wanna play?


(Provide -v flag to see the picked numbers.)
`
)

func verboseMode() {

	if len(os.Args[1:]) < 1 {
		fmt.Printf(usage, maxGuess)
		return
	}
	var verbose bool
	var guess int
	var err error
	if strings.ToLower(os.Args[1]) == "-v" {
		verbose = true
	}
	if verbose {
		guess, err = strconv.Atoi(os.Args[2])
	} else {
		guess, err = strconv.Atoi(os.Args[1])
	}

	if err != nil || guess < 0 {
		fmt.Println("Enter a Positive Integer Value")
		return
	}
	rand.Seed(time.Now().UnixNano())

	for x := 0; x <= maxGuess; x++ {
		n := rand.Intn(guess) + 1
		if verbose {
			fmt.Printf("%d ", n)
		}
		if n != guess {
			continue
		}
		if x == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!!")
		} else {
			fmt.Println("🎉  YOU WON!")
		}
		return
	}
	fmt.Println("☠️  YOU LOST... Try again?")
}
