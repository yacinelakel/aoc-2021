package day12

import (
	"fmt"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(lines []string) {
	start := state{"start", parse(&lines), []string{"start"}, false}
	fmt.Println(countPaths(start))
}

type state struct {
	n            string
	edges        edges
	visited      []string
	visitedTwice bool
}

func countPaths(s state) (int, int) {
	if s.n == "end" {
		return 1, 1
	}
	sum1, sum2 := 0, 0
	for _, adj := range s.edges[s.n] {
		nSum1, nSum2 := 0, 0
		switch {
		case strings.ToUpper(adj) == adj:
			nSum1, nSum2 = countPaths(state{adj, s.edges, s.visited, s.visitedTwice})
		case !common.Contains(adj, s.visited):
			nSum1, nSum2 = countPaths(state{adj, s.edges, append(s.visited, adj), s.visitedTwice})
		case !s.visitedTwice && adj != "start":
			_, nSum2 = countPaths(state{adj, s.edges, s.visited, !s.visitedTwice})
		}
		sum1, sum2 = sum1+nSum1, sum2+nSum2
	}
	return sum1, sum2
}

type edges map[string][]string

func (e *edges) addEdge(from, to string) {
	if fEdges, ok := (*e)[from]; !ok {
		(*e)[from] = []string{to}
	} else {
		(*e)[from] = append(fEdges, to)
	}
}

func parse(lines *[]string) edges {
	e := edges{}
	for _, line := range *lines {
		parts := strings.Split(line, "-")
		n0, n1 := parts[0], parts[1]
		e.addEdge(n0, n1)
		e.addEdge(n1, n0)
	}
	return e
}
