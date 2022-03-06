package section8

import "fmt"

func printTheRunes() {
	const word = "console"
	for _, v := range word {
		fmt.Printf("%c\n", v)
		fmt.Printf("\tdecimal: %[1]d\n", v)
		fmt.Printf("\thex    : 0x%[1]x\n", v)
		fmt.Printf("\tbinary : 0b%08[1]b\n", v)
	}

	fmt.Printf("with runes       : %s\n", string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))
	fmt.Printf("with decimals    : %s\n", string([]byte{99, 111, 110, 115, 111, 108, 101}))
	fmt.Printf("with hexadecimals: %s\n", string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
