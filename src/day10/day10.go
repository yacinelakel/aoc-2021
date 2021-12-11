package day10

import (
	"fmt"
	"sort"
	"strings"
)

func Run(lines []string) {
	fmt.Println(solve(lines))
}

func solve(lines []string) (int, int) {
	corruptedScore, incompleteScores := 0, []int{}
	for _, line := range lines {
		var pStack stack = []string{}
		corrupted := false
		for _, p := range strings.Split(line, "") {
			if strings.Contains(OPEN_PARENS, p) {
				pStack.push(p)
				continue
			}
			if strings.Contains(CLOSE_PARENS, p) {
				if !pStack.empty() && PAREN_MAP[pStack.pop()] == p {
					continue
				}
				corruptedScore += CORRUPTED_SCORE_MAP[p]
				corrupted = true
				break
			}
		}
		if !corrupted {
			incompleteScore := 0
			for !pStack.empty() {
				incompleteScore = (incompleteScore * 5) + INCOMPLETE_SCORE_MAP[pStack.pop()]
			}
			incompleteScores = append(incompleteScores, incompleteScore)
		}
	}
	sort.Ints(incompleteScores)
	return corruptedScore, incompleteScores[len(incompleteScores)/2]
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
	CORRUPTED_SCORE_MAP = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	INCOMPLETE_SCORE_MAP = map[string]int{
		"(": 1,
		"[": 2,
		"{": 3,
		"<": 4,
	}
)
