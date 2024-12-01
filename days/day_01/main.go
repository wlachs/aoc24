package day_01

import (
	"fmt"
	"github.com/wlchs/aoc24/utils"
	"sort"
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

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	list1 := make([]int, 0, len(input))
	list2 := make([]int, 0, len(input))

	for _, line := range input {
		raw := strings.Split(line, "   ")
		list1 = append(list1, utils.Atoi(raw[0]))
		list2 = append(list2, utils.Atoi(raw[1]))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for i := range input {
		sum += utils.Abs(list1[i] - list2[i])
	}

	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
