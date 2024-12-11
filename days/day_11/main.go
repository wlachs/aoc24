package day_11

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"math"
	"strconv"
	"strings"
)

// blink does the stone-splitting game logic
func blink(stones []uint64) []uint64 {
	var newStones []uint64
	for _, stone := range stones {
		if s1, ok1 := rule1(stone); ok1 {
			newStones = append(newStones, s1)
		} else if s2, ok2 := rule2(stone); ok2 {
			newStones = append(newStones, s2...)
		} else {
			newStones = append(newStones, rule3(stone))
		}
	}
	return newStones
}

// rule1 checks the first rule: stone = 0
func rule1(stone uint64) (uint64, bool) {
	if stone == 0 {
		return 1, true
	}
	return 0, false
}

// rule2 checks the first rule: #digit % 2 = 0
func rule2(stone uint64) ([]uint64, bool) {
	d := digits(stone)
	if d%2 == 0 {
		exp := uint64(math.Pow10(d / 2))
		return []uint64{stone / exp, stone % exp}, true
	}
	return nil, false
}

// digits counts the digits of the given uint64
func digits(number uint64) int {
	count := 0
	for number > 0 {
		number /= 10
		count++
	}
	return count
}

// rule3 checks the first rule: anything else
func rule3(stone uint64) uint64 {
	return stone * 2024
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
	stones := utils.ToUInt64Slice(strings.Split(input[0], " "))

	for range 25 {
		stones = blink(stones)
	}
	return strconv.Itoa(len(stones))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
