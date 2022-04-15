package section12

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item  item
	genre string
}

func queryById() {
	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		}, {
			item:  item{id: 2, name: "x-com 2", price: 30},
			genre: "stratergy",
		}, {
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
	}

	byID := make(map[int]game)
	for _, g := range games {
		byID[g.item.id] = g
	}

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
		> list   : lists all the games
		> id N   : queries a game by id
		> quit   : quits
  `)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.item.id, g.item.name, "("+g.genre+")", g.item.price)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.item.id, g.item.name, "("+g.genre+")", g.item.price)
		}
	}
}
