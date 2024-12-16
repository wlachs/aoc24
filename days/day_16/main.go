package day_16

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"math"
	"strconv"
)

// vecDir holds a position vector and a facing
type vecDir struct {
	vec, dir types.Vec2
}

// start gets the starting point of the map
func start(m map[types.Vec2]int32) types.Vec2 {
	for vec2, i := range m {
		if i == 'S' {
			return vec2
		}
	}
	panic("start not found")
}

// end gets the ending point of the map
func end(m map[types.Vec2]int32) types.Vec2 {
	for vec2, i := range m {
		if i == 'E' {
			return vec2
		}
	}
	panic("end not found")
}

// ways calculates the cost to travel from the starting node to every reachable node of the map
func ways(m map[types.Vec2]int32) map[vecDir]int {
	valid := func(vec types.Vec2) bool {
		return m[vec] != '#'
	}
	s := vecDir{start(m), types.Vec2{X: 1}}
	r := map[vecDir]int{}
	r[s] = 0
	visit := []vecDir{s}
	for len(visit) > 0 {
		cur := visit[0]
		visit = visit[1:]
		next := vecDir{cur.vec.Add(&cur.dir), cur.dir}
		left := vecDir{cur.vec, cur.dir.RotateLeft()}
		right := vecDir{cur.vec, cur.dir.RotateRight()}
		if n, ok := r[next]; valid(next.vec) && (!ok || r[cur]+1 < n) {
			r[next] = r[cur] + 1
			visit = append(visit, next)
		}
		if n, ok := r[left]; !ok || r[cur]+1000 < n {
			r[left] = r[cur] + 1000
			visit = append(visit, left)
		}
		if n, ok := r[right]; !ok || r[cur]+1000 < n {
			r[right] = r[cur] + 1000
			visit = append(visit, right)
		}
	}
	return r
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
	costs := ways(m)
	minCost := math.MaxInt
	e := end(m)
	for key, cost := range costs {
		if key.vec == e {
			minCost = min(minCost, cost)
		}
	}
	return strconv.Itoa(minCost)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
