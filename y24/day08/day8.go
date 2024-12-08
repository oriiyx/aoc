package day08

import (
	"fmt"

	"github.com/oriiyx/aoc/internal/data"
)

type Position struct {
	x, y int
}

type CharacterMap map[rune][]Position

func Day8Part1() {
	lines, _ := data.ReadLines("day8.txt")
	characters := make(CharacterMap)
	parse(lines, characters)
	r := transverse(characters, len(lines)-1, len(lines[0])-1)
	fmt.Println(r)
}

func Day8Part2() {
	lines, _ := data.ReadLines("day8.txt")
	characters := make(CharacterMap)
	parse(lines, characters)
	r := transverse2(characters, len(lines)-1, len(lines[0])-1)
	fmt.Println(r)
}

func transverse(
	characters CharacterMap,
	maxX, maxY int) int {
	antinodes := make(map[Position]bool)

	for _, positions := range characters {
		for i := 0; i < len(positions); i++ {
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}

				fPos := positions[i]
				sPos := positions[j]
				x := (fPos.x - sPos.x) + fPos.x
				y := (fPos.y - sPos.y) + fPos.y
				position := Position{x, y}

				if isValidPosition(position, maxX, maxY) && !antinodes[position] {
					antinodes[position] = true
				}
			}
		}
	}

	return len(antinodes)
}

func transverse2(
	characters CharacterMap,
	maxX, maxY int) int {
	antinodes := make(map[Position]bool)

	for _, positions := range characters {
		if len(positions) < 2 {
			continue
		}

		// First, mark all antenna positions as antinodes
		// (since they're in line with other antennas)
		for _, pos := range positions {
			antinodes[pos] = true
		}

		for i := 0; i < len(positions); i++ {
			for j := 0; j < len(positions); j++ {
				if i == j {
					continue
				}

				fPos := positions[i]
				sPos := positions[j]
				traverseX := fPos.x - sPos.x
				traverseY := fPos.y - sPos.y

				for {
					x := traverseX + fPos.x
					y := traverseY + fPos.y
					position := Position{x, y}

					// means we're out of the map and we need to stop walking
					if !isValidPosition(position, maxX, maxY) {
						break
					}

					fPos = Position{x, y}
					if !antinodes[position] {
						antinodes[position] = true
					}
				}

			}
		}
	}

	return len(antinodes)
}

func parse(lines []string, characters CharacterMap) {
	for y, line := range lines {
		for x, char := range line {
			pos := Position{x, y}
			switch char {
			case '.':
				continue
			default:
				characters[char] = append(characters[char], pos)
			}
		}
	}
}

func isValidPosition(pos Position, maxX, maxY int) bool {
	return pos.x >= 0 && pos.x <= maxX && pos.y >= 0 && pos.y <= maxY
}
