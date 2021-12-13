package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/yacinelakel/aoc-2021/common"
	"github.com/yacinelakel/aoc-2021/day1"
	"github.com/yacinelakel/aoc-2021/day10"
	"github.com/yacinelakel/aoc-2021/day11"
	"github.com/yacinelakel/aoc-2021/day12"
	"github.com/yacinelakel/aoc-2021/day13"
	"github.com/yacinelakel/aoc-2021/day2"
	"github.com/yacinelakel/aoc-2021/day3"
	"github.com/yacinelakel/aoc-2021/day4"
	"github.com/yacinelakel/aoc-2021/day5"
	"github.com/yacinelakel/aoc-2021/day6"
	"github.com/yacinelakel/aoc-2021/day7"
	"github.com/yacinelakel/aoc-2021/day8"
	"github.com/yacinelakel/aoc-2021/day9"
)

var DayRunners = map[int]func([]string){
	1:  day1.Run,
	2:  day2.Run,
	3:  day3.Run,
	4:  day4.Run,
	5:  day5.Run,
	6:  day6.Run,
	7:  day7.Run,
	8:  day8.Run,
	9:  day9.Run,
	10: day10.Run,
	11: day11.Run,
	12: day12.Run,
	13: day13.Run,
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Missing day argument")
		return
	}

	dayArg, err := strconv.Atoi(os.Args[1])
	common.FailOnErrors(err)

	if runner, ok := DayRunners[dayArg]; ok {
		runner(common.SplitNewLine(common.GetFileContent(dayArg)))
	} else {
		fmt.Printf("Day %d not implemented", dayArg)
	}
}
