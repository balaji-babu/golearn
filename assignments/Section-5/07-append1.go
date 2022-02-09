package section5

import (
	"bytes"
	"fmt"
)

func append1() {
	png, header := []byte{'P', 'N', 'G'}, []byte{}
	header = append(header, 'P', 'N', 'G')

	if bytes.Equal(png, header) {
		fmt.Println("They are equal.")
	}
}
