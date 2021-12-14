package day14

import (
	"fmt"
	"math"
	"strings"
)

func Run(lines []string) {
	template, rules := parse(lines)
	pm := pairMap{}
	fmt.Println(solve(template, rules, pm, 10), solve(template, rules, pm, 40))
}

func solve(template []string, rules ruleMap, pm pairMap, steps int) uint64 {
	resultMap := countMap{}
	for _, v := range template {
		resultMap[v] += uint64(1)
	}
	for i := 1; i < len(template); i++ {
		resultMap.add(countBetween(template[i-1], template[i], steps, &rules, &pm))
	}
	return result(&resultMap)
}

func result(c *countMap) uint64 {
	var min, max uint64
	min, max = math.MaxInt64, uint64(0)
	for _, v := range *c {
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}
	return max - min
}

func countBetween(left, right string, step int, rules *ruleMap, pm *pairMap) countMap {
	mid := rules.get(left + right)
	if step == 1 {
		return countMap{mid: 1}
	}
	if mem, found := pm.get(left+right, step); found {
		return mem
	}
	countMap := countMap{mid: uint64(1)}
	countMap.add(countBetween(left, mid, step-1, rules, pm))
	countMap.add(countBetween(mid, right, step-1, rules, pm))
	pm.set(left+right, step, countMap)
	return countMap
}

func parse(lines []string) ([]string, ruleMap) {
	template := strings.Split(lines[0], "")
	rules := ruleMap{}
	for _, l := range lines[2:] {
		parts := strings.Split(l, " -> ")
		rules[parts[0]] = parts[1]
	}
	return template, rules
}

type ruleMap map[string]string

func (r *ruleMap) get(p string) string {
	return (*r)[p]
}

type countMap map[string]uint64

func (m1 *countMap) add(m2 map[string]uint64) {
	for s, b2 := range m2 {
		(*m1)[s] += b2
	}
}

type countAtStepMap map[int](countMap)

type pairMap map[string](countAtStepMap)

func (pm *pairMap) set(p string, s int, c countMap) {
	if _, ok := (*pm)[p]; !ok {
		(*pm)[p] = countAtStepMap{s: c}
	} else {
		(*pm)[p][s] = c
	}
}

func (pm *pairMap) get(p string, i int) (countMap, bool) {
	if _, ok := (*pm)[p]; !ok {
		return nil, false
	}
	if _, ok := (*pm)[p][i]; !ok {
		return nil, false
	}
	return (*pm)[p][i], true
}
