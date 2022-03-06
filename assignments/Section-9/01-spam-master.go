package section9

import (
	"fmt"
	"os"
)

var (
	text = os.Args[1]
	size = len(text)
	buf  = []byte(text)
	in   bool
)

const (
	link  = "http://"
	nlink = len(link)
	mask  = '*'
)

func spamMaster() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide input to mask")
		return
	}

	fmt.Println("text size: ", size)
	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true

			i += nlink
		}
		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			buf[i] = mask
		}

	}
	fmt.Println(string(buf))
}
