package section12

import (
	"bufio"
	"encoding/json"
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

type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func dencode() {
	var games []game
	var decoded []jsonGame
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println("Sorry, there is a problem:", err)
		return
	}

	// load the data into usual game values
	for _, ds := range decoded {
		games = append(games, game{item{ds.ID, ds.Name, ds.Price}, ds.Genre})
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
  > save   : exports the data to json and quits
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
		case "save":

			// load the data into the encodable game values
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.item.id, g.item.name, g.genre, g.item.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}

			fmt.Println(string(out))
			return
		}
	}
}
