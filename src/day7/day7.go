package day7

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run() {
	input := common.SliceAtoi(strings.Split(common.GetFileContent(7), ","))
	sort.Ints(input)

	minFuel1, minFuel2 := math.MaxInt, math.MaxInt
	for p1 := input[0]; p1 < input[len(input)-1]; p1++ {
		fuel1, fuel2 := 0, 0
		for _, p2 := range input {
			d := int(math.Abs(float64(p2 - p1)))
			fuel1 += d
			fuel2 += d * (d + 1) / 2
		}
		switch {
		case minFuel1 > fuel1:
			minFuel1 = fuel1
		case minFuel2 > fuel2:
			minFuel2 = fuel2
		}
	}

	fmt.Println(minFuel1, minFuel2)
}
