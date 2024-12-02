package day_02

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"math"
	"strconv"
	"strings"
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

// isValid validates a report
func isValid(report []int) bool {
	valid := true
	sign := math.Signbit(float64(report[1] - report[0]))

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		absDiff := utils.Abs(diff)

		if math.Signbit(float64(diff)) != sign || absDiff < 1 || absDiff > 3 {
			valid = false
		}
	}

	return valid
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	count := 0

	for _, report := range input {
		values := utils.ToIntSlice(strings.Split(report, " "))

		if isValid(values) {
			count++
		}
	}

	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	count := 0

	for _, report := range input {
		values := utils.ToIntSlice(strings.Split(report, " "))

		if isValid(values) {
			count++
		} else {
			for i := range len(values) {
				adjustedValues := make([]int, 0, len(values)-1)
				adjustedValues = append(adjustedValues, values[:i]...)
				adjustedValues = append(adjustedValues, values[i+1:]...)

				if isValid(adjustedValues) {
					count++
					break
				}
			}
		}
	}

	return strconv.Itoa(count)
}
