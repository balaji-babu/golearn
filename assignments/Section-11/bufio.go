package section11

import (
	"bufio"
	"fmt"
	"os"
)

func bufioScanner() {
	defer os.Stdin.Close()
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		fmt.Println("scanned text: ", in.Text())
	}
	if err := in.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
