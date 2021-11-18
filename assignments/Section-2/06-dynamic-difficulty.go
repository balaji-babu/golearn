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

func dynamicDifficulty() {

	if len(os.Args[1:]) < 1 {
		fmt.Printf(usage, maxGuess)
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess < 0 {
		fmt.Println("Enter a Positive Integer Value")
		return
	}
	min := 10
	if guess > min {
		min = guess
	}
	rand.Seed(time.Now().UnixNano())

	for x := maxGuess + guess/4; x > 0; x-- {
		n := rand.Intn(guess) + 1
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
