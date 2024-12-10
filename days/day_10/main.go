package day_10

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"slices"
	"strconv"
)

// calculateScore calculates the score of a trailhead starting from the given position
func calculateScore(m map[types.Vec2]int32, vec types.Vec2) []types.Vec2 {
	if m[vec] == '9' {
		return []types.Vec2{vec}
	}

	var score []types.Vec2
	for _, nextVec := range vec.Around() {
		if m[nextVec]-m[vec] == 1 {
			for _, v := range calculateScore(m, nextVec) {
				if !slices.Contains(score, v) {
					score = append(score, v)
				}
			}
		}
	}

	return score
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
	m := utils.ParseInputToMap(input)
	score := 0

	for vec := range m {
		if m[vec] == '0' {
			score += len(calculateScore(m, vec))
		}
	}

	return strconv.Itoa(score)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
