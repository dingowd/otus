package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	const Msg = "Hello, OTUS!"
	revertMsg := stringutil.Reverse(Msg)
	fmt.Println(revertMsg)
}
