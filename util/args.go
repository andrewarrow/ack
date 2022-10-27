package util

import "os"

func GetArg(index int) string {
	if len(os.Args) > index {
		return os.Args[index]
	}
	return ""
}
