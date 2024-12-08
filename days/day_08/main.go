package day_08

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

	var uniqueAntiNodes []types.Vec2

	for _, positions := range frequencyVectorSliceMap {
		for i, n1 := range positions {
			for _, n2 := range positions[i+1:] {
				delta := n1.Subtract(&n2)
				antiNode1 := n1.Add(&delta)
				antiNode2 := n2.Subtract(&delta)

				if !slices.Contains(uniqueAntiNodes, antiNode1) {
					uniqueAntiNodes = append(uniqueAntiNodes, antiNode1)
				}

				if !slices.Contains(uniqueAntiNodes, antiNode2) {
					uniqueAntiNodes = append(uniqueAntiNodes, antiNode2)
				}
			}
		}
	}

	count := 0
	for _, node := range uniqueAntiNodes {
		if node.X <= bounds.X && node.Y <= bounds.Y && node.X >= 0 && node.Y >= 0 {
			count++
		}
	}

	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
