package day11

import (
	"fmt"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(lines []string) {
	grid := parse(&lines)
	fmt.Println(solve(grid))
}

func solve(grid [][]int) (int, int) {
	step, totalFlashes, simStep := 0, 0, 0
	for step < 100 || simStep == 0 {
		flashes := 0
		// Step 1
		for i := range grid {
			for j := range grid[i] {
				grid[i][j] += 1
			}
		}
		// Step 2
		for i := range grid {
			for j := range grid[i] {
				flashes += checkFlash(pos{i, j}, &grid)
			}
		}
		// Check if we have a result
		if step < 100 {
			totalFlashes += flashes
		}
		step++
		if simStep == 0 && flashes == 100 {
			simStep = step
		}
	}
	return totalFlashes, simStep
}

type pos struct {
	i int
	j int
}

func checkFlash(p pos, grid *[][]int) int {
	if (*grid)[p.i][p.j] <= 9 {
		return 0
	}
	(*grid)[p.i][p.j] = 0
	sum := 1
	for _, a := range getAdj(p) {
		if (*grid)[a.i][a.j] != 0 {
			(*grid)[a.i][a.j] += 1
			sum += checkFlash(a, grid)
		}
	}
	return sum
}

func getAdj(p pos) []pos {
	validAdj := []pos{}
	for _, a := range []pos{{p.i - 1, p.j}, {p.i + 1, p.j}, {p.i, p.j - 1}, {p.i, p.j + 1}, {p.i + 1, p.j + 1}, {p.i + 1, p.j - 1}, {p.i - 1, p.j - 1}, {p.i - 1, p.j + 1}} {
		if a.i < 0 || a.j < 0 || a.i > 9 || a.j > 9 {
			continue
		}
		validAdj = append(validAdj, a)
	}
	return validAdj
}

func parse(lines *[]string) [][]int {
	grid := [][]int{}
	for _, line := range *lines {
		row := common.SliceAtoi(strings.Split(line, ""))
		grid = append(grid, row)
	}
	return grid
}
