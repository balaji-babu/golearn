package switchStatement

import (
	"fmt"
	"os"
	"strconv"
)

func richerScale1() {
	if len(os.Args) == 2 {
		scale := os.Args[1]
		rsv, err := strconv.ParseFloat(scale, 32)
		if err != nil {
			fmt.Println("Not a valid input, please enter valid richer scale value")
			return
		}
		switch {
		case rsv <= 0.5:
			fmt.Printf("%.2f is micro", rsv)
		case rsv <= 2.5:
			fmt.Printf("%.2f is very minor", rsv)
		case rsv <= 3:
			fmt.Printf("%.2f is minor", rsv)
		case rsv <= 4.5:
			fmt.Printf("%.2f is light", rsv)
		case rsv <= 5:
			fmt.Printf("%.2f is moderate", rsv)
		case rsv <= 6:
			fmt.Printf("%.2f is strong", rsv)
		case rsv <= 7:
			fmt.Printf("%.2f is major", rsv)
		case rsv < 11:
			fmt.Printf("%.2f is great", rsv)
		default:
			fmt.Printf("%.2f is massive", rsv)

		}
	} else {
		fmt.Println("Enter the magnitude of earthquake")
		return
	}

}
