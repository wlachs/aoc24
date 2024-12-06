package day_06

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"slices"
	"strconv"
)

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
	var guardPosition types.Vec2
	guardDir := types.Vec2{X: 0, Y: -1}

	for pos, char := range m {
		if char == '^' {
			guardPosition = pos
			break
		}
	}

	var visitedLocations []types.Vec2
	for guardPosition.X >= 0 && guardPosition.X < len(input[0]) && guardPosition.Y >= 0 && guardPosition.Y < len(input) {
		if !slices.Contains(visitedLocations, guardPosition) {
			visitedLocations = append(visitedLocations, guardPosition)
		}

		nextPosition := guardPosition.Add(&guardDir)
		if m[nextPosition] != '#' {
			guardPosition = nextPosition
		} else {
			guardDir = guardDir.RotateRight()
		}
	}

	return strconv.Itoa(len(visitedLocations))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
