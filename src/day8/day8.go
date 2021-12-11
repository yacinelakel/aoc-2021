package day8

import (
	"fmt"
	"math"
	"regexp"
)

func Run(lines []string) {
	parsed := parse(&lines)
	fmt.Println(partOne(parsed), partTwo(parsed))
}

func partOne(entries []entry) int {
	count := 0
	for _, e := range entries {
		for _, o := range e.output {
			l := len(o)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				count++
			}
		}
	}
	return count
}

func partTwo(entries []entry) int {
	ans := 0
	for _, e := range entries {
		digitMap := map[string]int{}

		// Setup map of digits that are easy to determine
		one := find(2, e.input)[0]
		four := find(4, e.input)[0]
		seven := find(3, e.input)[0]
		eight := find(7, e.input)[0]
		digitMap[one] = 1
		digitMap[four] = 4
		digitMap[seven] = 7
		digitMap[eight] = 8

		// Determine digits with a segment length og 5
		for _, ei := range find(5, e.input) {
			if contains(ei, one) {
				digitMap[ei] = 3
			} else if contains(ei, remove(four, one)) {
				digitMap[ei] = 5
			} else {
				digitMap[ei] = 2
			}
		}

		// Determine digits with a segment length of 6
		for _, ei := range find(6, e.input) {
			if !contains(ei, one) {
				digitMap[ei] = 6
			} else if contains(ei, four) {
				digitMap[ei] = 9
			} else {
				digitMap[ei] = 0
			}
		}

		// Determine output
		for i, eo := range e.output {
			for _, ei := range e.input {
				if len(ei) == len(eo) && contains(ei, eo) {
					ans += digitMap[ei] * int(math.Pow10(len(e.output)-i-1))
				}
			}
		}
	}
	return ans
}

type entry struct {
	input  []string
	output []string
}

func find(l int, c []string) []string {
	res := []string{}
	for _, v := range c {
		if len(v) == l {
			res = append(res, v)
		}
	}
	return res
}

func remove(s1 string, s2 string) string {
	s := ""
	for _, r := range s1 {
		if !containsRune(r, s2) {
			s += string(r)
		}
	}
	return s
}

func contains(s1 string, s2 string) bool {
	res := true
	for _, r := range s2 {
		res = res && containsRune(r, s1)
	}
	return res
}

func containsRune(r rune, s string) bool {
	for _, rs := range s {
		if r == rs {
			return true
		}
	}
	return false
}

func parse(input *[]string) []entry {
	entries := []entry{}
	for _, l := range *input {
		pipeDil := regexp.MustCompile(`\s\|\s`)
		spaceDil := regexp.MustCompile(`\s`)
		parts := pipeDil.Split(l, -1)
		input := spaceDil.Split(parts[0], -1)
		output := spaceDil.Split(parts[1], -1)
		entries = append(entries, entry{input, output})
	}
	return entries
}
