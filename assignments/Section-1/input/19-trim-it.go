package trimSpace

import (
	"fmt"
	"strings"
)

func trimSpace() {
	msg := `
	The weather looks good.
	I should go and play




	`

	fmt.Println(strings.TrimSpace(msg))
}
