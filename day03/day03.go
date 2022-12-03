package day03

import (
	"fmt"
)

func Part1(in []string) (int, error) {
	var totalPriority int

	for _, l := range in {
		m := len(l) / 2
		l, r := l[:m], l[m:]

		p, err := overlapPriority(l, r)
		if err != nil {
			return 0, err
		}
		totalPriority += p
	}

	return totalPriority, nil
}

func Part2(in []string) (int, error) {
	var totalPriority int

	for i := 0; i < len(in); i += 3 {
		p, err := overlapPriority(in[i], in[i+1], in[i+2])
		if err != nil {
			return 0, err
		}
		totalPriority += p
	}

	return totalPriority, nil
}

func overlapPriority(backpacks ...string) (int, error) {
	var keep map[rune]struct{}
	for i, b := range backpacks {
		keepNext := make(map[rune]struct{})
		for _, x := range b {
			if _, ok := keep[x]; i == 0 || ok {
				keepNext[x] = struct{}{}
			}
		}
		keep = keepNext
	}

	if len(keep) != 1 {
		return 0, fmt.Errorf("day03: got overlap of %d items, wanted 1", len(keep))
	}

	for k := range keep {
		return priority(k)
	}

	panic("not reached")
}

func priority(x rune) (int, error) {
	switch {
	case 'a' <= x && x <= 'z':
		return int(x) - 'a' + 1, nil
	case 'A' <= x && x <= 'Z':
		return int(x) - 'A' + 27, nil
	default:
		return 0, fmt.Errorf("day03: invalid item: %c", x)
	}
}
