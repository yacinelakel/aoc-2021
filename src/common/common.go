package common

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %fs", name, elapsed.Seconds())
}

func SliceAtoi(sList []string) []int {
	iList := []int{}

	for _, s := range sList {
		if s == "" {
			continue
		}
		i, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		}
		iList = append(iList, i)
	}
	return iList
}

func Contains(s string, slice []string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}
	return false
}
