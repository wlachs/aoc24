package day_07

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"slices"
	"strconv"
	"strings"
)

// valid checks if the desired result can be achieved using various operators between the given values
func valid(result float64, values []float64) bool {
	if len(values) == 0 {
		return result == 0
	}
	return valid(result-values[0], values[1:]) || valid(result/values[0], values[1:])
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
	sum := 0

	for _, line := range input {
		split := strings.Split(line, ": ")
		result, _ := strconv.ParseFloat(split[0], 64)
		values := utils.ToFloatSlice(strings.Split(split[1], " "))

		slices.Reverse(values)

		if valid(result, values) {
			sum += int(result)
		}
	}
	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
