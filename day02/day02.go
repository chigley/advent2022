package day02

import (
	"fmt"
)

func Part1(in []Move) int {
	var score int
	for _, m := range in {
		score += m.Score()
	}
	return score
}

type Move struct {
	Oppt, Us rune
}

func (m *Move) Score() int {
	oppt := int(m.Oppt - 'A')
	us := int(m.Us - 'X')

	score := us + 1

	switch us {
	case oppt:
		score += 3
	case (oppt + 1) % 3:
		score += 6
	}

	return score
}

func Parse(in []string) ([]Move, error) {
	moves := make([]Move, len(in))

	for i, l := range in {
		if _, err := fmt.Sscanf(l, "%c %c", &moves[i].Oppt, &moves[i].Us); err != nil {
			return nil, fmt.Errorf("day02: parsing %q: %w", l, err)
		}
	}

	return moves, nil
}
