package day_15

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"slices"
	"strconv"
)

// direction enum
var (
	UP    = types.Vec2{Y: -1}
	RIGHT = types.Vec2{X: 1}
	DOWN  = types.Vec2{Y: 1}
	LEFT  = types.Vec2{X: -1}
)

// readMapAndInstructions parses the input box map and instruction set
func readMapAndInstructions(input []string) (map[types.Vec2]int32, []types.Vec2) {
	n := slices.Index(input, "")
	m := utils.ParseInputToMap(input[:n])
	var instructions []types.Vec2
	for _, row := range input[n+1:] {
		for _, c := range row {
			switch c {
			case '^':
				instructions = append(instructions, UP)
			case '>':
				instructions = append(instructions, RIGHT)
			case 'v':
				instructions = append(instructions, DOWN)
			case '<':
				instructions = append(instructions, LEFT)
			}
		}
	}
	return m, instructions
}

// robot gets the robot's current positions
func robot(m map[types.Vec2]int32) types.Vec2 {
	for vec2, c := range m {
		if c == '@' {
			return vec2
		}
	}
	panic("robot not found")
}

// push tries to recursively push the item at the given position in the given direction
func push(m map[types.Vec2]int32, pos types.Vec2, dir *types.Vec2, save bool) bool {
	cascade := func(side types.Vec2) bool {
		if push(m, pos.Add(dir), dir, false) && push(m, pos.Add(dir).Add(&side), dir, false) {
			if save {
				push(m, pos.Add(dir), dir, save)
				push(m, pos.Add(dir).Add(&side), dir, save)
				m[pos.Add(dir)], m[pos.Add(dir).Add(&side)] = m[pos], m[pos.Add(&side)]
				m[pos], m[pos.Add(&side)] = '.', '.'
			}
			return true
		}
		return false
	}
	switch m[pos] {
	case '[':
		if dir.Y != 0 {
			return cascade(RIGHT)
		}
		fallthrough
	case ']':
		if dir.Y != 0 {
			return cascade(LEFT)
		}
		fallthrough
	case '@':
		fallthrough
	case 'O':
		if push(m, pos.Add(dir), dir, save) {
			m[pos.Add(dir)] = m[pos]
			m[pos] = '.'
			return true
		}
	case '.':
		return true
	}
	return false
}

// gps calculates the "GPS coordinates" of a box
func gps(vec2 types.Vec2) int {
	return vec2.X + 100*vec2.Y
}

// sum calculates the sum of box GPS coordinates
func sum(m map[types.Vec2]int32) int {
	s := 0
	for vec2, c := range m {
		if c == 'O' || c == '[' {
			s += gps(vec2)
		}
	}
	return s
}

// extend modifies the input map by doubling the width of every element other than '@', then returns the new map
func extend(m map[types.Vec2]int32) map[types.Vec2]int32 {
	extended := map[types.Vec2]int32{}
	for vec2, c := range m {
		switch c {
		case '@':
			extended[types.Vec2{X: vec2.X * 2, Y: vec2.Y}], extended[types.Vec2{X: vec2.X*2 + 1, Y: vec2.Y}] = c, '.'
		case 'O':
			extended[types.Vec2{X: vec2.X * 2, Y: vec2.Y}], extended[types.Vec2{X: vec2.X*2 + 1, Y: vec2.Y}] = '[', ']'
		default:
			extended[types.Vec2{X: vec2.X * 2, Y: vec2.Y}], extended[types.Vec2{X: vec2.X*2 + 1, Y: vec2.Y}] = c, c
		}
	}
	return extended
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
	m, instructions := readMapAndInstructions(input)
	for _, instruction := range instructions {
		pos := robot(m)
		push(m, pos, &instruction, true)
	}
	return strconv.Itoa(sum(m))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	m, instructions := readMapAndInstructions(input)
	m = extend(m)
	for _, instruction := range instructions {
		pos := robot(m)
		push(m, pos, &instruction, true)
	}
	return strconv.Itoa(sum(m))
}
