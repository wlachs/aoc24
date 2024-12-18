package day_18

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"math"
	"strconv"
	"strings"
)

// bytes reads the input and gets the falling bytes in a slice of vectors
func bytes(input []string) []types.Vec2 {
	fallingBytes := make([]types.Vec2, len(input))
	for i, s := range input {
		coords := strings.Split(s, ",")
		fallingBytes[i] = types.Vec2{X: utils.Atoi(coords[0]), Y: utils.Atoi(coords[1])}
	}
	return fallingBytes
}

// initMap initializes the map without any obstacles
func initMap(size int) map[types.Vec2]int32 {
	m := map[types.Vec2]int32{}
	for y := range size {
		for x := range size {
			m[types.Vec2{X: x, Y: y}] = '.'
		}
	}
	return m
}

// dijkstra finds the shortest path between the two positions if possible
func dijkstra(m map[types.Vec2]int32, size int) int {
	costMap := map[types.Vec2]int{}
	for vec2 := range m {
		costMap[vec2] = math.MaxInt
	}
	s := types.Vec2{}
	e := types.Vec2{X: size - 1, Y: size - 1}
	costMap[s] = 0
	visit := []types.Vec2{s}
	for len(visit) > 0 {
		h := visit[0]
		visit = visit[1:]
		for _, next := range h.Around() {
			if c, ok := m[next]; ok && c != '#' && costMap[h]+1 < costMap[next] {
				visit = append(visit, next)
				costMap[next] = costMap[h] + 1
			}
		}
	}
	return costMap[e]
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input, 71, 1024))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input, 71))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string, size, steps int) string {
	fallingBytes := bytes(input)
	m := initMap(size)
	for i := range steps {
		m[fallingBytes[i]] = '#'
	}
	return strconv.Itoa(dijkstra(m, size))
}

// Part2 solves the second part of the exercise
func Part2(input []string, size int) string {
	fallingBytes := bytes(input)
	m := initMap(size)
	for i := range len(input) {
		m[fallingBytes[i]] = '#'
		if dijkstra(m, size) == math.MaxInt {
			return input[i]
		}
	}
	panic("no solution")
}
