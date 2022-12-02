package day02_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2022"
	"github.com/chigley/advent2022/day02"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 15, 12},
	{"input", 12794, 14979},
}

func TestDay02(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2022.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			moves, err := day02.Parse(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day02.Part1(moves))
			assert.Equal(t, tt.part2, day02.Part2(moves))
		})
	}
}
