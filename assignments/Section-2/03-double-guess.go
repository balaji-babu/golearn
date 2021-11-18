package luckyGuess

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
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
`
)

func doubleGuess() {
	if len(os.Args[1:]) < 1 {
		fmt.Printf(usage, maxGuess)
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Println("Enter a Positive Integer Value")
		return
	}
	guess2, err := strconv.Atoi(os.Args[2])
	if err != nil || guess2 < 0 {
		fmt.Println("Enter a Positive Integer Value")
		return
	}
	min := guess
	if guess < guess2 {
		min = guess2
	}
	rand.Seed(time.Now().UnixNano())

	for x := 0; x <= maxGuess; x++ {
		n := rand.Intn(min) + 1

		if n == guess || n == guess2 {
			fmt.Println("🎉  YOU WIN!")
			return
		}
	}

	fmt.Println("☠️  YOU LOST... Try again?")
}
