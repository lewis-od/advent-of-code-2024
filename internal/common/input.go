package common

import (
	"errors"
	"os"
)

var ErrInvalidInvocation = errors.New("expcted 1 argument")

func ReadInputFile() []byte {
	if len(os.Args) != 2 {
		panic(ErrInvalidInvocation)
	}
	filename := os.Args[1]
	return Must(os.ReadFile(filename))
}
