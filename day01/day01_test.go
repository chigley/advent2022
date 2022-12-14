package day01_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2022"
	"github.com/chigley/advent2022/day01"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 24000, 45000},
	{"input", 70764, 203905},
}

func TestDay01(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2022.ReadIntGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day01.Part1(in))
			assert.Equal(t, tt.part2, day01.Part2(in))
		})
	}
}
