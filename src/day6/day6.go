package day6

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/yacinelakel/aoc-2021/common"
)

func Run(lines []string) {
	fish := common.SliceAtoi(strings.Split(lines[0], ","))
	mem := map[int]*big.Int{}
	fmt.Println(solve(80, fish, &mem), solve(256, fish, &mem))
}

func solve(days int, fish []int, mem *map[int]*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, f := range fish {
		sum.Add(sum, gen(days-f, mem))
	}
	return sum.Add(sum, big.NewInt(int64(len(fish))))
}

func gen(day int, mem *map[int]*big.Int) *big.Int {
	if day <= 0 {
		return big.NewInt(0)
	}
	if val, ok := (*mem)[day]; ok {
		return val
	}
	sum := big.NewInt(1)
	sum.Add(sum, (gen(day-9, mem)))
	sum.Add(sum, (gen(day-7, mem)))
	(*mem)[day] = sum
	return sum
}
