package day5

import (
	"fmt"
	"regexp"

	"github.com/yacinelakel/aoc-2021/common"
)

type pos struct {
	x int
	y int
}

type seg struct {
	p1 pos
	p2 pos
}

func Run() {
	input := common.SplitNewLine(common.GetFileContent(5))
	fmt.Println(solve(parseInput(input)))
}

func solve(segs []seg) (int, int) {
	m1, m2 := map[pos]int{}, map[pos]int{}
	for _, s := range segs {
		p := s.p1
		mark(s, p, &m1, &m2)
		for p != s.p2 {
			p = pos{move(p.x, s.p2.x), move(p.y, s.p2.y)}
			mark(s, p, &m1, &m2)
		}
	}

	return count(&m1), count(&m2)
}

func count(m *map[pos]int) int {
	c := 0
	for _, h := range *m {
		if h > 1 {
			c++
		}
	}
	return c
}

func mark(s seg, p pos, m1 *map[pos]int, m2 *map[pos]int) {
	if s.p1.x == s.p2.x || s.p1.y == s.p2.y {
		(*m1)[p]++
	}
	(*m2)[p]++
}

func move(a, b int) int {
	if a < b {
		return a + 1
	} else if a > b {
		return a - 1
	}
	return a
}

func parseInput(lines []string) []seg {
	segs := []seg{}
	for _, l := range lines {
		re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
		parts := common.SliceAtoi(re.FindStringSubmatch(l)[1:])
		s := seg{pos{parts[0], parts[1]}, pos{parts[2], parts[3]}}
		segs = append(segs, s)
	}
	return segs
}
