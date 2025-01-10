package utils

import "github.com/wlchs/aoc24/types"

// FindUniqueInMap finds a unique c value in a map
func FindUniqueInMap(m map[types.Vec2]int32, c int32) (types.Vec2, bool) {
	for vec2, char := range m {
		if char == c {
			return vec2, true
		}
	}
	return types.Vec2{}, false
}
