package day03

import "fmt"

func Part1(in []string) (int, error) {
	var totalPriority int
	for _, l := range in {
		p, err := overlapPriority(l)
		if err != nil {
			return 0, err
		}
		totalPriority += p
	}
	return totalPriority, nil
}

func overlapPriority(backpack string) (int, error) {
	m := len(backpack) / 2
	l, r := backpack[:m], backpack[m:]

	seen := make(map[rune]struct{})
	for _, x := range l {
		seen[x] = struct{}{}
	}

	for _, x := range r {
		if _, ok := seen[x]; ok {
			return priority(x)
		}
	}

	return 0, fmt.Errorf("day03: no overlap in backpack: %q", backpack)
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
