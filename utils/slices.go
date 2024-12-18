package utils

import "strconv"

// ToIntSlice converts a string slice to an int slice
func ToIntSlice(numbers []string) []int {
	s := make([]int, 0, len(numbers))
	for _, number := range numbers {
		s = append(s, Atoi(number))
	}
	return s
}

// ToUInt8Slice converts a string slice to an uint8 slice
func ToUInt8Slice(numbers []string) []uint8 {
	s := make([]uint8, len(numbers))
	for i, u := range ToUInt64Slice(numbers) {
		s[i] = uint8(u)
	}
	return s
}

// ToUInt64Slice converts a string slice to an uint64 slice
func ToUInt64Slice(numbers []string) []uint64 {
	s := make([]uint64, 0, len(numbers))
	for _, number := range numbers {
		i, _ := strconv.ParseUint(number, 10, 64)
		s = append(s, i)
	}
	return s
}

// ToFloatSlice converts a string slice to a float64 slice
func ToFloatSlice(numbers []string) []float64 {
	s := make([]float64, 0, len(numbers))
	for _, number := range numbers {
		f, _ := strconv.ParseFloat(number, 64)
		s = append(s, f)
	}
	return s
}

// ToStringSlice converts an int slice to a string slice
func ToStringSlice(numbers []int) []string {
	s := make([]string, 0, len(numbers))
	for _, number := range numbers {
		s = append(s, strconv.Itoa(number))
	}
	return s
}

// EqualsUInt8Slice compares two uint8 slices
func EqualsUInt8Slice(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// PrefixUInt8Slice checks if slice b is a prefix of slice a
func PrefixUInt8Slice(a, b []uint8) bool {
	if len(a) < len(b) {
		return false
	}
	for i := range b {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
