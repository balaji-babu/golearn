package trimRight

import (
	"fmt"
	"strings"
)

func trimRight() {
	name := "Balaji           "
	fmt.Println(len(strings.TrimRight(name, " ")))
}
