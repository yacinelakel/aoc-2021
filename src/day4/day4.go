package day4

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(lines []string) {
	numList, boards := parse(lines)
	fmt.Println(solve(numList, boards))
}

func solve(numbers []int, boards [][][]int) (int, int) {
	boardsWithBingo := map[int]bool{}
	var firstBingo, lastBingo int
	for _, n := range numbers {
		for bi, b := range boards {
			if boardsWithBingo[bi] {
				continue
			}

			p, e := findNumber(b, n)

			if e != nil {
				continue
			}

			b[p.i][p.j] = -1

			if isBingo(b, p) {
				if len(boardsWithBingo) == 0 {
					firstBingo = calculateSum(b, n)
				}
				boardsWithBingo[bi] = true
				if len(boards) == len(boardsWithBingo) {
					lastBingo = calculateSum(b, n)
				}
			}
		}
	}
	return firstBingo, lastBingo
}

func calculateSum(board [][]int, num int) int {
	sum := 0
	for _, r := range board {
		for _, v := range r {
			if v != -1 {
				sum += v
			}
		}
	}
	return sum * num
}

func isBingo(board [][]int, p pos) bool {
	if board[p.i][0] == -1 && board[p.i][1] == -1 && board[p.i][2] == -1 && board[p.i][3] == -1 && board[p.i][4] == -1 {
		return true
	}
	if board[0][p.j] == -1 && board[1][p.j] == -1 && board[2][p.j] == -1 && board[3][p.j] == -1 && board[4][p.j] == -1 {
		return true
	}
	return false
}

func findNumber(board [][]int, num int) (pos, error) {
	for i, r := range board {
		for j, v := range r {
			if v == num {
				return pos{i, j}, nil
			}
		}
	}
	return pos{}, errors.New("not found")
}

func parse(input []string) ([]int, [][][]int) {
	lines := input
	numList := common.SliceAtoi(strings.Split(lines[0], ","))
	boards := [][][]int{}
	for i := 2; i < len(lines); i++ {
		if (i-1)%6 == 0 {
			continue
		}
		bI := (i - 2) / 6
		rI := (i - 2) % 6
		if rI == 0 {
			boards = append(boards, [][]int{})
		}
		row := common.SliceAtoi(strings.Split(lines[i], " "))
		boards[bI] = append(boards[bI], []int{})
		boards[bI][rI] = append(boards[bI][rI], row...)
	}
	return numList, boards
}

type pos struct {
	i int
	j int
}
