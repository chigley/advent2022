package day01

import (
	"github.com/chigley/advent2022"
)

func Part1(in [][]int) int {
	sums := make([]int, len(in))
	for i, ns := range in {
		sums[i] = advent2022.Sum(ns)
	}

	return advent2022.MaxInts(sums)
}
