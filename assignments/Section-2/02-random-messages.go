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

func randomNumberGenerator() {

	if len(os.Args[1:]) != 1 {
		fmt.Printf(usage, maxGuess)
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Println("Enter a Positive Integer Value")
		return
	}
	rand.Seed(time.Now().UnixNano())

	for x := 0; x <= maxGuess; x++ {
		n := rand.Intn(guess) + 1
		if n == guess {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉  YOU WIN!")
			case 1:
				fmt.Println("🎉  YOU'RE AWESOME!")
			case 2:
				fmt.Println("🎉  PERFECT!")
			}
			return
		}
	}

	msg := "%s Try again?\n"

	switch rand.Intn(2) {
	case 0:
		fmt.Printf(msg, "☠️  YOU LOST...")
	case 1:
		fmt.Printf(msg, "☠️  JUST A BAD LUCK...")
	}
}
