package day1

import (
	"fmt"
	"strconv"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run() {
	raw := common.GetFileContent(1)
	input := toIntArray(common.SplitNewLine(raw))

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func toIntArray(strArr []string) []int {
	var intArr = make([]int, len(strArr))
	for i := range intArr {
		d, e := strconv.Atoi(strArr[i])
		common.FailOnErrors(e)
		intArr[i] = d
	}
	return intArr
}

func partOne(input []int) int {
	countIncreased := 0
	for i := 1; i < len(input); i++ {
		if input[i-1] < input[i] {
			countIncreased++
		}
	}
	return countIncreased
}

func partTwo(input []int) int {
	countIncreased := 0
	for i := 3; i < len(input); i++ {
		a, b, c, d := input[i-3], input[i-2], input[i-1], input[i]
		windowA := a + b + c
		windowB := b + c + d
		if windowA < windowB {
			countIncreased++
		}
	}
	return countIncreased
}
