package advent2022

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func readStrings(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2022: scanner: %w", err)
	}

	return lines, nil
}

func readStringGroups(r io.Reader) ([][]string, error) {
	lines, err := readStrings(r)
	if err != nil {
		return nil, fmt.Errorf("advent2022: reading strings: %w", err)
	}

	var groups [][]string

	var group []string
	for _, l := range lines {
		if l == "" {
			groups = append(groups, group)
			group = nil
			continue
		}

		group = append(group, l)
	}

	groups = append(groups, group)

	return groups, nil
}

func ReadIntGroups(name string) ([][]int, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("advent2022: opening %q: %w", name, err)
	}
	defer f.Close()

	return readIntGroups(f)
}

func readIntGroups(r io.Reader) ([][]int, error) {
	stringGroups, err := readStringGroups(r)
	if err != nil {
		return nil, fmt.Errorf("advent2022: reading string groups: %w", err)
	}

	intGroups := make([][]int, len(stringGroups))
	for i, g := range stringGroups {
		intGroups[i] = make([]int, len(g))

		for j, s := range g {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, fmt.Errorf("advent2022: atoi: %w", err)
			}

			intGroups[i][j] = n
		}
	}

	return intGroups, nil
}
