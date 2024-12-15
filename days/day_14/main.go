package day_14

import (
	"fmt"
	"github.com/wlchs/aoc24/types"
	"github.com/wlchs/aoc24/utils"
	"regexp"
	"strconv"
)

// robot struct containing position and velocity
type robot struct {
	p, v types.Vec2
}

// quadrant returns the ID of the quadrant where the robot current is
// If the robot is exactly on the border between two quadrants, return false
func (r robot) quadrant(width, height int) (int, bool) {
	w, h := width/2, height/2
	if r.p.X == w || r.p.Y == h {
		return -1, false
	}
	return 2*(r.p.X/(w+1)) + r.p.Y/(h+1), true
}

// parseRobots parses the input and returns a slice of robots
func parseRobots(input []string) []robot {
	robots := make([]robot, 0, len(input))
	re := regexp.MustCompile(`.*=(?P<pX>\d+),(?P<pY>\d+).*=(?P<vX>.*),(?P<vY>.*)`)
	for _, s := range input {
		match := re.FindStringSubmatch(s)
		p := types.Vec2{X: utils.Atoi(match[1]), Y: utils.Atoi(match[2])}
		v := types.Vec2{X: utils.Atoi(match[3]), Y: utils.Atoi(match[4])}
		robots = append(robots, robot{p, v})
	}
	return robots
}

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input, 101, 103))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string, width, height int) string {
	robots := parseRobots(input)

	// printRobots(robots, width, height)
	for range 100 {
		for i, r := range robots {
			robots[i].p = types.Vec2{X: utils.Mod(r.p.X+r.v.X, width), Y: utils.Mod(r.p.Y+r.v.Y, height)}
		}
		//printRobots(robots, width, height)
	}
	quadrants := make([]uint64, 4)
	for _, r := range robots {
		if q, ok := r.quadrant(width, height); ok {
			quadrants[q]++
		}
	}
	for _, r := range robots {
		q, ok := r.quadrant(width, height)
		fmt.Printf("%v -> %d(%v)\n", r, q, ok)
	}
	for i, quadrant := range quadrants {
		fmt.Printf("%d->%d\n", i, quadrant)
	}
	return strconv.FormatUint(quadrants[0]*quadrants[1]*quadrants[2]*quadrants[3], 10)
}

func printRobots(robots []robot, width, height int) {
	m := map[types.Vec2]int{}
	for _, r := range robots {
		m[r.p]++
	}
	for y := range height {
		for x := range width {
			i, ok := m[types.Vec2{X: x, Y: y}]
			if ok {
				fmt.Print(i)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
