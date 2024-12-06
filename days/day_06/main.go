package day_06

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"maps"
	"slices"
	"strconv"
)

// pair represents a 2D position and a facing
type pair struct {
	pos types.Vec2
	dir types.Vec2
}

// findVisitedPositions counts how many positions the guard visits before leaving the map and returns the list of visited coordinates as well
// if the guard is stuck in a loop, return -1
func findVisitedPositions(guardPosition types.Vec2, guardDir types.Vec2, m map[types.Vec2]int32) (int, []types.Vec2) {
	var visitedLocations []types.Vec2
	var history []pair

	for ok := true; ok; _, ok = m[guardPosition] {

		if !slices.Contains(visitedLocations, guardPosition) {
			visitedLocations = append(visitedLocations, guardPosition)
		}

		p := pair{pos: guardPosition, dir: guardDir}
		if !slices.Contains(history, p) {
			history = append(history, p)
		} else {
			return -1, visitedLocations
		}

		nextPosition := guardPosition.Add(&guardDir)

		if m[nextPosition] != '#' {
			guardPosition = nextPosition
		} else {
			guardDir = guardDir.RotateRight()
		}
	}

	return len(visitedLocations), visitedLocations
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
	var guardPosition types.Vec2

	for pos, char := range m {
		if char == '^' {
			guardPosition = pos
			break
		}
	}

	visitedLocations, _ := findVisitedPositions(guardPosition, types.Vec2{X: 0, Y: -1}, m)
	return strconv.Itoa(visitedLocations)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	var guardPosition types.Vec2

	for pos, char := range m {
		if char == '^' {
			guardPosition = pos
			break
		}
	}

	loops := 0
	_, path := findVisitedPositions(guardPosition, types.Vec2{X: 0, Y: -1}, m)

	for _, toReplace := range path {
		if m[toReplace] != '.' {
			continue
		}

		clone := maps.Clone(m)
		clone[toReplace] = '#'

		if count, _ := findVisitedPositions(guardPosition, types.Vec2{X: 0, Y: -1}, clone); count == -1 {
			loops++
		}
	}

	return strconv.Itoa(loops)
}
