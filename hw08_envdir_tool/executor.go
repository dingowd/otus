package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	extProg := cmd[0]
	args := cmd[1:]
	ext := exec.Command(extProg, args...)
	ext.Stdout = os.Stdout
	ext.Stdin = os.Stdin
	ext.Stderr = os.Stderr
	for key, val := range env {
		if val.NeedRemove {
			os.Unsetenv(key)
		} else {
			os.Unsetenv(key)
			os.Setenv(key, val.Value)
		}
	}
	ext.Env = os.Environ()
	err := ext.Run()
	returnCode = 0
	if err != nil {
		returnCode = 1
	}
	return
}
