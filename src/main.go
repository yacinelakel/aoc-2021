package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yacinelakel/aoc-2021/common"
	"github.com/yacinelakel/aoc-2021/day1"
	"github.com/yacinelakel/aoc-2021/day2"
	"github.com/yacinelakel/aoc-2021/day3"
	"github.com/yacinelakel/aoc-2021/day4"
	"github.com/yacinelakel/aoc-2021/day5"
	"github.com/yacinelakel/aoc-2021/day6"
	"github.com/yacinelakel/aoc-2021/day7"
	"github.com/yacinelakel/aoc-2021/day8"
	"github.com/yacinelakel/aoc-2021/day9"
)

func main() {
	dayArg, err := strconv.Atoi(os.Args[1])
	common.FailOnErrors(err)

	getDayFunc(dayArg)()
}

func getDayFunc(day int) func() {
	dayMap := map[int]func(){
		1: day1.Run,
		2: day2.Run,
		3: day3.Run,
		4: day4.Run,
		5: day5.Run,
		6: day6.Run,
		7: day7.Run,
		8: day8.Run,
		9: day9.Run,
	}

	if runner, ok := dayMap[day]; ok {
		return runner
	} else {
		return func() { fmt.Printf("Day %d not implemented", day) }
	}
}
