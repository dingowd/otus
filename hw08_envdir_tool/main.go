package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	env, err := ReadDir(args[0])
	if err != nil {
		log.Fatal(err)
	}
	exitCode := RunCmd(args[1:], env)
	os.Exit(exitCode)
}
