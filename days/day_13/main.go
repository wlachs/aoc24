package day_13

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"regexp"
	"strconv"
)

// scenario describes a test case
type scenario struct {
	A, B, X types.Vec2
}

// parseInput receives the input and creates a scenario slice from it
func parseInput(input []string, advanced bool) []scenario {
	res := make([]scenario, 0, (len(input)+1)/4)
	cur := scenario{}
	re := regexp.MustCompile(".*X=?(?P<X>.*), Y=?(?P<Y>.*)")
	for i, s := range input {
		match := re.FindStringSubmatch(s)
		switch i % 4 {
		case 0:
			cur.A = types.Vec2{X: utils.Atoi(match[1]), Y: utils.Atoi(match[2])}
		case 1:
			cur.B = types.Vec2{X: utils.Atoi(match[1]), Y: utils.Atoi(match[2])}
		case 2:
			if advanced {
				cur.X = types.Vec2{X: utils.Atoi(match[1]) + 10000000000000, Y: utils.Atoi(match[2]) + 10000000000000}
			} else {
				cur.X = types.Vec2{X: utils.Atoi(match[1]), Y: utils.Atoi(match[2])}
			}
			res = append(res, cur)
		}
	}
	return res
}

// minCost calculates the minimum cost of tickets required to get the price if possible
func minCost(s scenario) uint64 {
	a := float64(s.X.X*s.B.Y-s.X.Y*s.B.X) / float64(s.A.X*s.B.Y-s.A.Y*s.B.X)
	b := float64(s.X.Y*s.A.X-s.X.X*s.A.Y) / float64(s.A.X*s.B.Y-s.A.Y*s.B.X)
	if a == float64(uint64(a)) && b == float64(uint64(b)) {
		return 3*uint64(a) + uint64(b)
	}
	return 0
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	scenarios := parseInput(input, false)
	cost := uint64(0)
	for _, s := range scenarios {
		cost += minCost(s)
	}
	return strconv.FormatUint(cost, 10)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	scenarios := parseInput(input, true)
	cost := uint64(0)
	for _, s := range scenarios {
		cost += minCost(s)
	}
	return strconv.FormatUint(cost, 10)
}
