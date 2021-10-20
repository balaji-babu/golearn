package challenge

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOK = "Access granted to %q.\n"
)

func challenge02() {
	if len(os.Args) < 3 {
		fmt.Printf(usage)
	} else if (os.Args[1] == "jack" && os.Args[2] == "1888") || (os.Args[1] == "inanc" && os.Args[2] == "1879") {
		fmt.Printf(accessOK, os.Args[1])
	} else if (os.Args[1] != "jack") || (os.Args[1] != "inanc") {
		fmt.Printf(errUser, os.Args[1])
	} else {
		fmt.Printf(errPwd, os.Args[1])
	}
}
