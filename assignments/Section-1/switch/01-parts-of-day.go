package switchStatement

import (
	"fmt"
	"time"
)

func partsOfADay() {
	switch time := time.Now().Hour(); {
	case time >= 7 && time < 12:
		fmt.Println("Good Morning")
	case time >= 12 && time < 17:
		fmt.Println("Good Afternoon")
	case time >= 17 && time < 19:
		fmt.Println("Good Evening")
	default:
		fmt.Println("Good Night")
	}
}
