package day_20

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"math"
	"strconv"
)

// f implements the counter logic to find out how many different ways are available to cheat in the puzzle
func f(m map[types.Vec2]int32, distance, threshold int) int {
	p := dijkstra(m, s(m), e(m))
	count := 0
	for i, start := range p {
		for j, end := range p[i:] {
			wallPath := utils.Abs(end.X-start.X) + utils.Abs(end.Y-start.Y)
			if wallPath <= distance && j-wallPath >= threshold {
				count++
			}
		}
	}
	return count
}

// s finds the start position in the map
func s(m map[types.Vec2]int32) types.Vec2 {
	vec, _ := utils.FindUniqueInMap(m, 'S')
	return vec
}

// e finds the end position in the map
func e(m map[types.Vec2]int32) types.Vec2 {
	vec, _ := utils.FindUniqueInMap(m, 'E')
	return vec
}

// dijkstra finds the shortest path between the two positions if possible
func dijkstra(m map[types.Vec2]int32, start, end types.Vec2) []types.Vec2 {
	costMap := map[types.Vec2]int{}
	for vec2 := range m {
		costMap[vec2] = math.MaxInt
	}
	costMap[start] = 0
	visit := []types.Vec2{start}
	for len(visit) > 0 {
		h := visit[0]
		visit = visit[1:]
		for _, next := range h.Around() {
			if c, ok := m[next]; ok && c != '#' && costMap[h]+1 < costMap[next] {
				visit = append(visit, next)
				costMap[next] = costMap[h] + 1
				if next == end {
					return path(costMap, start, end)
				}
			}
		}
	}
	panic("no path found")
}

// path finds one of the shortest paths on the map from the start to the end position
func path(costMap map[types.Vec2]int, start types.Vec2, end types.Vec2) []types.Vec2 {
	cur := end
	p := []types.Vec2{cur}
	for cur != start {
		for _, vec2 := range cur.Around() {
			if costMap[vec2] < costMap[cur] {
				cur = vec2
			}
		}
		p = append([]types.Vec2{cur}, p...)
	}
	return p
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input, 100))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input, 100))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string, threshold int) string {
	m := utils.ParseInputToMap(input)
	return strconv.Itoa(f(m, 2, threshold))
}

// Part2 solves the second part of the exercise
func Part2(input []string, threshold int) string {
	m := utils.ParseInputToMap(input)
	return strconv.Itoa(f(m, 20, threshold))
}
