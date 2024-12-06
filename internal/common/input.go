package common

import (
	"errors"
	"os"
	"strings"
)

var ErrInvalidInvocation = errors.New("expcted 1 argument")

func ReadInputFileLines() []string {
	contents := ReadInputFile()
	return strings.Split(string(contents), "\n")
}

func ReadInputFile() []byte {
	if len(os.Args) != 2 {
		panic(ErrInvalidInvocation)
	}
	filename := os.Args[1]
	return Must(os.ReadFile(filename))
}
