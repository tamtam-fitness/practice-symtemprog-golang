package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer    = strings.NewReader("COMPUTER")
	system      = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader

	a := io.NewSectionReader(programming, 5, 1)
	s := io.LimitReader(system, 1)
	c := io.LimitReader(computer, 1)
	i := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(a, s, c, i)
	io.Copy(os.Stdout, stream)
}
