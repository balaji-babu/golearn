package rawConcat

import (
	"fmt"
	"os"
)

func rawConcat() {
	name := `hi` + os.Args[1] + `!
	how are you?
	`
	fmt.Printf("%s", name)
}
