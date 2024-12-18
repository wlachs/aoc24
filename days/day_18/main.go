package day_18

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"math"
	"slices"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input, 71, 1024))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string, size, steps int) string {
	fallingBytes := make([]types.Vec2, len(input))
	for i, s := range input {
		coords := strings.Split(s, ",")
		fallingBytes[i] = types.Vec2{X: utils.Atoi(coords[0]), Y: utils.Atoi(coords[1])}
	}
	fmt.Println(fallingBytes)
	m := map[types.Vec2]int32{}
	for y := range size {
		for x := range size {
			m[types.Vec2{X: x, Y: y}] = '.'
		}
	}
	for i := range steps {
		m[fallingBytes[i]] = '#'
	}
	m[types.Vec2{}] = 'S'
	m[types.Vec2{X: size - 1, Y: size - 1}] = 'E'
	printMap(m, size)
	path := dijkstra(m, types.Vec2{}, types.Vec2{X: size - 1, Y: size - 1})
	printMap(m, size, path...)
	return strconv.Itoa(len(path) - 1)
}

// dijkstra finds the shortest path between the two positions if possible
func dijkstra(m map[types.Vec2]int32, s types.Vec2, e types.Vec2) []types.Vec2 {
	costMap := map[types.Vec2]int{}
	for vec2 := range m {
		costMap[vec2] = math.MaxInt
	}
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
	path := []types.Vec2{e}
	cur := e
	for cur != s {
		leastCost := math.MaxInt
		var leastCostVec types.Vec2
		for _, v := range cur.Around() {
			if c, ok := costMap[v]; ok && c < leastCost {
				leastCost = c
				leastCostVec = v
			}
		}
		path = append(path, leastCostVec)
		cur = leastCostVec
	}
	return path
}

func printMap(m map[types.Vec2]int32, size int, path ...types.Vec2) {
	for y := range size {
		for x := range size {
			v := types.Vec2{X: x, Y: y}
			if slices.Contains(path, v) {
				fmt.Print("O")
			} else {
				fmt.Printf("%s", string(m[v]))
			}
		}
		fmt.Println()
	}
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
