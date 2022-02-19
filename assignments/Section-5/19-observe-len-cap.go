package section5

import "fmt"

func olc() {
	games := []string{}
	fmt.Printf("Length of games string %d\n", len(games))
	fmt.Printf("Capacity of games string %d\n", cap(games))

	games = append(games, "pacman", "mario", "tetris", "doom")
	fmt.Printf("New length of games string %d\n", len(games))
	fmt.Printf("New capacity of games string %d\n", cap(games))

	// --- #2 ---
	fmt.Println()

	for k := range games {
		n := games[:k]
		fmt.Printf("games[:%d]'s len: %d cap: %d\n", k, len(n), cap(n))
	}

	// --- #3 ---
	fmt.Println()
	zero := games[0:0]
	fmt.Printf("games's len: %d cap: %d\n", len(games), cap(games))
	fmt.Printf("zero len: %d cap: %d\n", len(zero), cap(zero))
	for _, v := range []string{"one", "two", "three", "four", "five"} {
		zero = append(zero, v)
		fmt.Printf("zero's len: %d cap: %d\n", len(games), cap(games))
	}

	// --- #4 ---
	fmt.Println()
	for k := range zero {
		n := zero[:k]
		fmt.Printf("zero[:%d]'s len: %d cap: %d\n", k, len(n), cap(n))
	}

	// --- #5 ---
	fmt.Println()
	zero = zero[:cap(zero)]
	for k := range zero {
		n := zero[:k]
		fmt.Printf("zero[:%d]'s len: %d cap: %d\n", k, len(n), cap(n))
	}

	// --- #6 ---
	fmt.Println()
	zero[1] = "new"
	games[1] = "new"
	fmt.Printf("zero  : %q\n", zero)
	fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// newCap := games[:cap(games)+1]
	// fmt.Printf("newCap: %v", newCap)
	// Output ==> Runtime as expected (panic: runtime error: slice bounds out of range [:5] with capacity 4)
}
