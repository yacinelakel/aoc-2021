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
)

func main() {
	dayArg, err := strconv.Atoi(os.Args[1])
	common.FailOnErrors(err)

	runDay(dayArg)()
}

func runDay(day int) func() {
	switch day {
	case 1:
		return day1.Run
	case 2:
		return day2.Run
	case 3:
		return day3.Run
	case 4:
		return day4.Run
	case 5:
		return day5.Run
	default:
		return (func() { fmt.Printf("No runner for day %d", day) })
	}
}
