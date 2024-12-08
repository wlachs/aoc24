package day_08

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"slices"
	"strconv"
)

// buildFrequencyVectorSliceMap parses the input map and finds on which positions which frequencies can be found
func buildFrequencyVectorSliceMap(m map[types.Vec2]int32) (map[int32][]types.Vec2, types.Vec2) {
	frequencyVectorSliceMap := map[int32][]types.Vec2{}
	bounds := types.Vec2{X: 0, Y: 0}

	for pos, frequency := range m {
		if pos.X > bounds.X || pos.Y > bounds.Y {
			bounds = pos
		}

		if frequency == '.' {
			continue
		}

		_, ok := frequencyVectorSliceMap[frequency]
		if ok {
			frequencyVectorSliceMap[frequency] = append(frequencyVectorSliceMap[frequency], pos)
		} else {
			frequencyVectorSliceMap[frequency] = []types.Vec2{pos}
		}
	}

	return frequencyVectorSliceMap, bounds
}

// addUniqueNodeWhileInBounds adds an anti-node on the equal distances from the given node
func addUniqueNodeWhileInBounds(uniqueAntiNodes []types.Vec2, pos types.Vec2, delta types.Vec2, bounds types.Vec2) []types.Vec2 {
	node := pos.Add(&delta)
	if node.X <= bounds.X && node.Y <= bounds.Y && node.X >= 0 && node.Y >= 0 && !slices.Contains(uniqueAntiNodes, node) {
		uniqueAntiNodes = append(uniqueAntiNodes, node)
	}

	return uniqueAntiNodes
}

// addUniqueNodesWhileInBounds adds every anti-node on the equal distances from the given node until the end of the map is reached
func addUniqueNodesWhileInBounds(uniqueAntiNodes []types.Vec2, pos types.Vec2, delta types.Vec2, bounds types.Vec2) []types.Vec2 {
	for node := pos; node.X <= bounds.X && node.Y <= bounds.Y && node.X >= 0 && node.Y >= 0; node = node.Add(&delta) {
		if !slices.Contains(uniqueAntiNodes, node) {
			uniqueAntiNodes = append(uniqueAntiNodes, node)
		}
	}

	return uniqueAntiNodes
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
	frequencyVectorSliceMap, bounds := buildFrequencyVectorSliceMap(m)

	var uniqueAntiNodes []types.Vec2
	for _, positions := range frequencyVectorSliceMap {
		for i, n1 := range positions {
			for _, n2 := range positions[i+1:] {
				delta := n1.Subtract(&n2)
				uniqueAntiNodes = addUniqueNodeWhileInBounds(uniqueAntiNodes, n1, delta, bounds)
				uniqueAntiNodes = addUniqueNodeWhileInBounds(uniqueAntiNodes, n2, delta.Multiply(-1), bounds)
			}
		}
	}

	return strconv.Itoa(len(uniqueAntiNodes))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	frequencyVectorSliceMap, bounds := buildFrequencyVectorSliceMap(m)

	var uniqueAntiNodes []types.Vec2
	for _, positions := range frequencyVectorSliceMap {
		for i, n1 := range positions {
			for _, n2 := range positions[i+1:] {
				delta := n1.Subtract(&n2)
				uniqueAntiNodes = addUniqueNodesWhileInBounds(uniqueAntiNodes, n1, delta, bounds)
				uniqueAntiNodes = addUniqueNodesWhileInBounds(uniqueAntiNodes, n2, delta.Multiply(-1), bounds)
			}
		}
	}

	return strconv.Itoa(len(uniqueAntiNodes))
}
