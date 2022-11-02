package util

import (
	"os"
	"strconv"
)

func GetArg(index int) string {
	if len(os.Args) > index {
		return os.Args[index]
	}
	return ""
}

func Atoi(num string, ifzero int) int {
	thing, _ := strconv.Atoi(num)
	if thing == 0 {
		return ifzero
	}
	return thing
}
