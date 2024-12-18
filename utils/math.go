package utils

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Mod is a Python-like implementation of the modulo function
func Mod(a, b int) int {
	return (a%b + b) % b
}
