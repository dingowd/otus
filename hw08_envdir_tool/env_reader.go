package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ParseValue parsing first string of file to set environment.
func ParseValue(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	fScanner := bufio.NewScanner(file)
	zero := []byte{0x00}
	fScanner.Scan()
	toProcess := fScanner.Bytes()
	toProcess = bytes.ReplaceAll(toProcess, zero, []byte("\n"))
	strToReturn := string(toProcess)
	// strToReturn = strings.Split(strToReturn, "\n")[0] // to cut after /n
	strToReturn = strings.TrimRight(strToReturn, "\t ")
	return strToReturn, nil
}

func Path(dir string) string {
	if string(dir[len(dir)-1]) != "/" {
		return dir + "/"
	}
	return dir
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	env := make(Environment)
	var toHelp EnvValue
	for _, fileProp := range fileInfo {
		if fileProp.Size() == 0 {
			toHelp.NeedRemove = true
			toHelp.Value = ""
		} else {
			toHelp.NeedRemove = false
			toHelp.Value, err = ParseValue(Path(dir) + fileProp.Name())
			if err != nil {
				return nil, err
			}
		}
		env[fileProp.Name()] = toHelp
	}
	return env, nil
}
