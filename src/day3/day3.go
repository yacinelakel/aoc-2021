package day3

import (
	"fmt"
	"math"
	"strconv"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run() {
	lines := common.SplitNewLine(common.GetFileContent(3))
	fmt.Println(partOne(lines))
	fmt.Println(partTwo(lines))
}

func partOne(lines []string) int {
	onesMap := make(map[int]int)

	for i := 0; i < len(lines); i++ {
		bits := lines[i]
		for j := 0; j < len(bits); j++ {
			if bits[j] == '1' {
				onesMap[j] = onesMap[j] + 1
			}
		}
	}

	gamma, eps := 0.0, 0.0
	for i := 0; i < len(lines[0]); i++ {
		if onesMap[i] > len(lines)/2 {
			gamma += (math.Pow(2, float64(len(lines[0])-i-1)))
		} else {
			eps += (math.Pow(2, float64(len(lines[0])-i-1)))
		}
	}
	return int(gamma * eps)
}

func partTwo(lines []string) int {
	oxy := applyCriteria(lines[:], func(list0, list1 []string) []string {
		if len(list1) >= len(list0) {
			return list1
		}
		return list0
	})

	co2 := applyCriteria(lines[:], func(list0, list1 []string) []string {
		if len(list0) <= len(list1) {
			return list0
		}
		return list1
	})

	return oxy * co2
}

func applyCriteria(list []string, critera func([]string, []string) []string) int {
	cList := list[:]
	for i := 0; i < len(list[0]); i++ {
		if len(cList) == 1 {
			break
		}
		list0, list1 := splitListAtIndex(cList, i)
		cList = critera(list1, list0)
	}
	res, e := strconv.ParseInt(cList[0], 2, 64)
	if e != nil {
		panic(e)
	}
	return int(res)
}

func splitListAtIndex(list []string, i int) ([]string, []string) {
	list0 := make([]string, 0)
	list1 := make([]string, 0)
	for _, l := range list {
		if l[i] == '1' {
			list1 = append(list1, l)
		} else {
			list0 = append(list0, l)
		}
	}
	return list0, list1
}
