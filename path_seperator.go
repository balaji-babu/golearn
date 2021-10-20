package main

import (
	"fmt"
	"path"
)

func Pathspliter() {

	dir, file := path.Split("c://Users/balaj/Documents/balaji/test.txt")

	fmt.Println("file:", file)
	fmt.Println("dir:", dir)
}
