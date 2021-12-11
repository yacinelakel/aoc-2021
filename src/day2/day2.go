package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(lines []string) {
	fmt.Println(partOne(lines), partTwo(lines))
}

type pos struct {
	h int
	v int
	a int
}

type command struct {
	dir  string
	unit int
}

func parseCommand(line string) command {
	sArr := strings.Split(line, " ")
	dir := sArr[0]
	num, err := strconv.Atoi(sArr[1])
	common.FailOnErrors(err)
	return command{dir, num}
}

func partOne(input []string) int {
	var p pos
	for _, l := range input {
		c := parseCommand(l)
		switch c.dir {
		case "forward":
			p.h += c.unit
		case "down":
			p.v += c.unit
		case "up":
			p.v -= c.unit
		default:
			panic("invalid direction")
		}
	}

	return p.h * p.v
}

func partTwo(input []string) int {
	var p pos
	for _, l := range input {
		c := parseCommand(l)
		switch c.dir {
		case "forward":
			p.h += c.unit
			p.v += (p.a * c.unit)
		case "down":
			p.a += c.unit
		case "up":
			p.a -= c.unit
		default:
			panic("invalid direction")
		}
	}

	return p.h * p.v
}
