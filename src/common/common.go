package common

import (
	"fmt"
	"os"
	"strings"
)

func FailOnErrors(params ...error) {
	if len(params) != 0 {
		for _, e := range params {
			if e != nil {
				panic(e)
			}
		}
	}
}

func GetFileContent(day int) string {
	fileLocation := fmt.Sprintf("./day%d/input.txt", day)

	data, err := os.ReadFile(fileLocation)
	FailOnErrors(err)

	return string(data)
}

func SplitNewLine(s string) []string {
	return strings.Split(s, "\r\n")
}
