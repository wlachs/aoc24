package day_05

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"slices"
	"strconv"
	"strings"
)

// checkAndCorrectInput tests whether the input is valid and if not, tries to correct it
func checkAndCorrectInput(values []int, constraints []types.Vec2) {
	for y := 0; y < len(values); y++ {
		for x := y + 1; x < len(values); x++ {
			if slices.Contains(constraints, types.Vec2{X: values[x], Y: values[y]}) {
				tmp := values[y]
				values[y] = values[x]
				values[x] = tmp
			}
		}
	}
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
	emptyLineId := slices.Index(input, "")
	constraints := make([]types.Vec2, 0, emptyLineId)
	sum := 0

	for _, row := range input[:emptyLineId] {
		values := utils.ToIntSlice(strings.Split(row, "|"))
		constraints = append(constraints, types.Vec2{X: values[0], Y: values[1]})
	}

	for _, row := range input[emptyLineId+1:] {
		values := utils.ToIntSlice(strings.Split(row, ","))
		clone := slices.Clone(values)

		checkAndCorrectInput(clone, constraints)

		if slices.Compare(values, clone) == 0 {
			sum += clone[(len(clone)-1)/2]
		}
	}

	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	emptyLineId := slices.Index(input, "")
	constraints := make([]types.Vec2, 0, emptyLineId)
	sum := 0

	for _, row := range input[:emptyLineId] {
		values := utils.ToIntSlice(strings.Split(row, "|"))
		constraints = append(constraints, types.Vec2{X: values[0], Y: values[1]})
	}

	for _, row := range input[emptyLineId+1:] {
		values := utils.ToIntSlice(strings.Split(row, ","))
		clone := slices.Clone(values)

		checkAndCorrectInput(clone, constraints)

		if slices.Compare(values, clone) != 0 {
			sum += clone[(len(clone)-1)/2]
		}
	}

	return strconv.Itoa(sum)
}
