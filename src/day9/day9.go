package day9

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(raw string) {
	board := parse(common.SplitNewLine(raw))
	fmt.Println(solve(board))
}

func solve(board [][]int) (int, int) {
	basins := []int{}
	partOne := 0
	for i, r := range board {
		for j, b := range r {
			p, isLowPt := pos{b, i, j}, true
			for _, a := range getAdj(p, board) {
				isLowPt = isLowPt && b < a.v
			}
			if !isLowPt {
				continue
			}
			partOne += b + 1
			size := bfs(p, board)
			basins = append(basins, size)
		}
	}
	sort.Ints(basins)
	partTwo := 1
	for i := len(basins) - 1; i > len(basins)-4; i-- {
		partTwo *= basins[i]
	}
	return partOne, partTwo
}

type pos struct {
	v int
	i int
	j int
}

func bfs(p pos, board [][]int) int {
	visited, queue, size := map[pos]bool{p: true}, []pos{p}, 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		size++
		for _, a := range getAdj(cur, board) {
			if !visited[a] && a.v != 9 {
				visited[a] = true
				queue = append(queue, a)
			}
		}
	}
	return size
}

func getAdj(p pos, board [][]int) []pos {
	aList := []pos{}
	if p.i != 0 {
		aList = append(aList, pos{board[p.i-1][p.j], p.i - 1, p.j})
	}
	if p.i != len(board)-1 {
		aList = append(aList, pos{board[p.i+1][p.j], p.i + 1, p.j})
	}
	if p.j != 0 {
		aList = append(aList, pos{board[p.i][p.j-1], p.i, p.j - 1})
	}
	if p.j != len((board)[p.i])-1 {
		aList = append(aList, pos{board[p.i][p.j+1], p.i, p.j + 1})
	}
	return aList
}

func parse(lines []string) [][]int {
	board := [][]int{}
	for _, l := range lines {
		row := []int{}
		for _, r := range l {
			s, e := strconv.Atoi(string(r))
			if e != nil {
				panic(e)
			}
			row = append(row, int(s))
		}
		board = append(board, row)
	}
	return board
}
