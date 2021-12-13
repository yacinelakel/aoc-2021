package day13

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(lines []string) {
	dots, instructions := parse(lines)
	partOne, partTwo := solve(dots, instructions)
	fmt.Println(partOne)
	printDots(partTwo)
}

func solve(dots []dot, instructions []instruction) (int, []dot) {
	partOne := 0
	for i, inst := range instructions {
		if inst.foldHorizontal {
			dots = foldHorizontal(dots, inst.value)
		} else {
			dots = foldVertical(dots, inst.value)
		}
		if i == 0 {
			partOne = len(dots)
		}
	}
	return partOne, dots
}

func foldHorizontal(dots []dot, foldP int) []dot {
	return fold(dots, foldP, func(d dot) int { return d.y }, func(d dot, v int) dot { return dot{d.x, v} })
}

func foldVertical(dots []dot, foldP int) []dot {
	return fold(dots, foldP, func(d dot) int { return d.x }, func(d dot, v int) dot { return dot{v, d.y} })
}

func fold(dots []dot, foldP int, posFn func(dot) int, newDotFn func(dot, int) dot) []dot {
	// Sort
	sort.Slice(dots, func(i, j int) bool {
		return posFn(dots[i]) < posFn(dots[j])
	})
	// split
	index := 0
	for posFn(dots[index]) < foldP {
		index++
	}
	part1, part2 := dots[0:index], dots[index:]
	// Remove overlapping dots
	for _, d := range part2 {
		nDot := newDotFn(d, int(math.Abs(float64(2*foldP-posFn(d)))))
		if !containsDot(&part1, nDot) {
			part1 = append(part1, nDot)
		}
	}
	return part1
}

func containsDot(dots *[]dot, d dot) bool {
	for _, e := range *dots {
		if e == d {
			return true
		}
	}
	return false
}

type instruction struct {
	foldHorizontal bool
	value          int
}

type dot struct {
	x int
	y int
}

func parse(lines []string) ([]dot, []instruction) {
	dots, instructions := []dot{}, []instruction{}
	i, line := 0, lines[0]
	for line != "" {
		parts := common.SliceAtoi(strings.Split(line, ","))
		dots = append(dots, dot{parts[0], parts[1]})
		i++
		line = lines[i]
	}
	i++
	for i < len(lines) {
		line = lines[i]
		parts := strings.Split(line, "=")
		v, e := strconv.Atoi(parts[1])
		if e != nil {
			panic(e)
		}
		if strings.Contains(parts[0], "x") {
			instructions = append(instructions, instruction{false, v})
		} else {
			instructions = append(instructions, instruction{true, v})
		}
		i++
	}
	return dots, instructions
}

func printDots(dots []dot) {
	sort.Slice(dots, func(i, j int) bool {
		if dots[i].y == dots[j].y {
			return dots[i].x < dots[j].x
		}
		return dots[i].y < dots[j].y
	})
	cursorX, cusrsorY := 0, 0
	for _, d := range dots {
		diffY := d.y - cusrsorY
		for i := 0; i < diffY; i++ {
			fmt.Println()
			cusrsorY, cursorX = d.y, 0
		}
		for i := cursorX; i < d.x; i++ {
			fmt.Print(" ")
		}
		fmt.Print("#")
		cursorX = d.x + 1
	}
}
