package day04_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2022"
	"github.com/chigley/advent2022/day04"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 2},
	{"input", 547},
}

func TestDay04(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2022.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			ranges, err := day04.Parse(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day04.Part1(ranges))
		})
	}
}
