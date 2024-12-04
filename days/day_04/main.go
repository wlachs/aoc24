package day_04

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"strconv"
)

// search recursively moves in the provided direction and tries to match the given string.
// If a match is found, the result is 1, otherwise 0
func search(m map[types.Vec2]int32, pos types.Vec2, s string, dir types.Vec2) int {
	if s == "" {
		return 1
	}

	head := int32(s[0])
	if m[pos] != head {
		return 0
	}

	return search(m, pos.Add(&dir), s[1:], dir)
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
	count := 0

	for y := range input {
		for x := range input[y] {
			head := types.Vec2{X: x, Y: y}
			count += search(m, head, "XMAS", types.Vec2{X: 0, Y: -1})
			count += search(m, head, "XMAS", types.Vec2{X: 1, Y: -1})
			count += search(m, head, "XMAS", types.Vec2{X: 1, Y: 0})
			count += search(m, head, "XMAS", types.Vec2{X: 1, Y: 1})
			count += search(m, head, "XMAS", types.Vec2{X: 0, Y: 1})
			count += search(m, head, "XMAS", types.Vec2{X: -1, Y: 1})
			count += search(m, head, "XMAS", types.Vec2{X: -1, Y: 0})
			count += search(m, head, "XMAS", types.Vec2{X: -1, Y: -1})
		}
	}

	return strconv.Itoa(count)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m := utils.ParseInputToMap(input)
	count := 0
	dir := types.Vec2{X: 1, Y: 1}

	for y := range input {
		for x := range input[y] {
			head := types.Vec2{X: x, Y: y}

			if (search(m, head, "AS", dir) != 1 || search(m, head, "AM", dir.Multiply(-1)) != 1) && (search(m, head, "AM", dir) != 1 || search(m, head, "AS", dir.Multiply(-1)) != 1) {
				continue
			}

			if (search(m, head, "AS", dir.RotateRight()) == 1 && search(m, head, "AM", dir.RotateRight().Multiply(-1)) == 1) || search(m, head, "AM", dir.RotateRight()) == 1 && search(m, head, "AS", dir.RotateRight().Multiply(-1)) == 1 {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}
