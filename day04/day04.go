package day04

import (
	"fmt"

	"github.com/chigley/advent2022"
)

func Part1(in [][2]advent2022.XY) int {
	var pairs int
	for _, p := range in {
		if fullyContains(p[0], p[1]) {
			pairs++
		}
	}
	return pairs
}

func fullyContains(a, b advent2022.XY) bool {
	return a.X <= b.X && a.Y >= b.Y || a.X >= b.X && a.Y <= b.Y
}

func Parse(in []string) ([][2]advent2022.XY, error) {
	ranges := make([][2]advent2022.XY, len(in))
	for i, l := range in {
		if _, err := fmt.Sscanf(l, "%d-%d,%d-%d", &ranges[i][0].X, &ranges[i][0].Y, &ranges[i][1].X, &ranges[i][1].Y); err != nil {
			return nil, err
		}
	}
	return ranges, nil
}
