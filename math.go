package advent2022

import "math"

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxInts(ns []int) int {
	if len(ns) == 0 {
		panic("max of empty slice")
	}

	max := math.MinInt
	for _, n := range ns {
		max = Max(max, n)
	}
	return max
}

func Sum(ns []int) int {
	var sum int

	for _, n := range ns {
		sum += n
	}

	return sum
}
