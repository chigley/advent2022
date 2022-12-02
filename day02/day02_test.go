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
}{
	{"example", 15},
	{"input", 12794},
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
		})
	}
}
