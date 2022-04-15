package section12

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item  item
	genre string
}

func list() {
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

	for _, gs := range games {
		fmt.Printf("#%d: %-15q %-20s $%d\n", gs.item.id, gs.item.name, "("+gs.genre+")", gs.item.price)
	}
}
