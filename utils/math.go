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

// Pow is a purely int-based pow function
func Pow(a, b int) int {
	r := 1
	for range b {
		r *= a
	}
	return r
}
