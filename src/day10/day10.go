package day10

import (
	"fmt"
	"sort"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(raw string) {
	input := common.SplitNewLine(raw)
	fmt.Println(solve(input))
}

func solve(lines []string) (int, int) {
	partOne := 0
	scores := []int{}
	for _, line := range lines {
		var s stack = []string{}
		valid := true
		for _, p := range strings.Split(line, "") {
			if strings.Contains(OPEN_PARENS, p) {
				s.push(p)
				continue
			}
			if strings.Contains(CLOSE_PARENS, p) {
				if !s.empty() {
					last := s.pop()
					if PAREN_MAP[last] == p {
						continue
					}
				}
				partOne += ERROR_SCORE_MAP[p]
				valid = false
				break
			}
		}
		if valid {
			score := 0
			for !s.empty() {
				score = (score * 5) + CLOSE_SCORE_MAP[s.pop()]
			}
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	partTwo := scores[len(scores)/2]
	return partOne, partTwo
}

type stack []string

func (s *stack) pop() string {
	l := len(*s)
	e := (*s)[l-1]
	*s = (*s)[:l-1]
	return e
}

func (s *stack) push(e string) {
	*s = append(*s, e)
}

func (s *stack) empty() bool {
	return len(*s) == 0
}

var (
	OPEN_PARENS  = "([{<"
	CLOSE_PARENS = ")]}>"
	PAREN_MAP    = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
		"<": ">",
	}
	ERROR_SCORE_MAP = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	CLOSE_SCORE_MAP = map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
)
