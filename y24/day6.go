package y24

import (
	"fmt"
	"log"

	"github.com/oriiyx/aoc/internal/data"
)

type Position struct {
	x, y int
}

type Direction struct {
	dx, dy int
	next   *Direction
}

type Guard struct {
	pos     Position
	dir     *Direction
	visited map[Position]bool
	counter int
}

type State struct {
	pos Position
	dir *Direction
}

func Day6Part1() {
	lines, err := data.ReadLines("day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	guard := &Guard{
		visited: make(map[Position]bool),
		dir:     createDirections(),
		counter: 1,
	}
	obstacles := make(map[Position]bool)

	parse(lines, guard, obstacles)
	walk(guard, obstacles, len(lines)-1, len(lines[0])-1)

	fmt.Println(guard.counter)
}

func createDirections() *Direction {
	up := &Direction{dx: 0, dy: -1}
	right := &Direction{dx: 1, dy: 0}
	down := &Direction{dx: 0, dy: 1}
	left := &Direction{dx: -1, dy: 0}

	up.next = right
	right.next = down
	down.next = left
	left.next = up

	return up
}

func parse(lines []string, guard *Guard, obstacles map[Position]bool) {
	for y, line := range lines {
		for x, char := range line {
			pos := Position{x, y}
			switch char {
			case '^':
				guard.pos = pos
				guard.visited[pos] = true
			case '#':
				obstacles[pos] = true
			}
		}
	}
}

func walk(guard *Guard, obstacles map[Position]bool, maxX, maxY int) {
	for {
		newPos := Position{
			x: guard.pos.x + guard.dir.dx,
			y: guard.pos.y + guard.dir.dy,
		}

		if !isValidPosition(newPos, maxX, maxY) {
			break
		}

		if obstacles[newPos] {
			guard.dir = guard.dir.next
			continue
		}

		guard.pos = newPos
		if !guard.visited[newPos] {
			guard.visited[newPos] = true
			guard.counter++
		}
	}
}

func isValidPosition(pos Position, maxX, maxY int) bool {
	return pos.x >= 0 && pos.x <= maxX && pos.y >= 0 && pos.y <= maxY
}

func Day6Part2() {
	lines, err := data.ReadLines("day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	guard := &Guard{
		visited: make(map[Position]bool),
		dir:     createDirections(),
		counter: 1,
	}
	obstacles := make(map[Position]bool)

	parse(lines, guard, obstacles)
	validPositions := findLoopPositions(guard, obstacles, len(lines)-1, len(lines[0])-1)

	fmt.Println(validPositions)
}

func findLoopPositions(guard *Guard, obstacles map[Position]bool, maxX, maxY int) int {
	validPositions := 0
	startPos := guard.pos

	// Try each position
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			pos := Position{x, y}
			if obstacles[pos] || pos == startPos {
				continue
			}

			obstacles[pos] = true

			if wouldCreateLoop(guard, obstacles, maxX, maxY) {
				validPositions++
			}

			// Remove test obstruction
			delete(obstacles, pos)
		}
	}

	return validPositions
}

func wouldCreateLoop(guard *Guard, obstacles map[Position]bool, maxX, maxY int) bool {
	visited := make(map[State]bool)
	current := State{pos: guard.pos, dir: guard.dir}

	for {
		newPos := Position{
			x: current.pos.x + current.dir.dx,
			y: current.pos.y + current.dir.dy,
		}

		if !isValidPosition(newPos, maxX, maxY) {
			return false // Not a loop if we hit boundary
		}

		if obstacles[newPos] {
			current.dir = current.dir.next
			continue
		}

		current.pos = newPos
		if visited[current] {
			return true
		}

		visited[current] = true
	}
}
