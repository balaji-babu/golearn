package toLowerCase

import (
	"fmt"
	"os"
	"strings"
)

func toLowerCase() {
	fmt.Println(strings.ToLower(os.Args[1]))
}
