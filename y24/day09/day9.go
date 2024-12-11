package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oriiyx/aoc/internal/data"
)

// Block represents a single block on the disk
type Block struct {
	FileID  int
	IsSpace bool
}

// parseDiskMap converts the input string into a sequence of file and space lengths
func parseDiskMap(input string) []int {
	result := make([]int, len(input))
	for i, ch := range input {
		num, _ := strconv.Atoi(string(ch))
		result[i] = num
	}
	return result
}

// createInitialBlocks creates the initial block representation with file IDs
func createInitialBlocks(lengths []int) []Block {
	var blocks []Block
	fileID := 0

	for i := 0; i < len(lengths); i += 2 {
		// Add file blocks
		for j := 0; j < lengths[i]; j++ {
			blocks = append(blocks, Block{FileID: fileID, IsSpace: false})
		}
		fileID++

		// Add space blocks if not at the end
		if i+1 < len(lengths) {
			for j := 0; j < lengths[i+1]; j++ {
				blocks = append(blocks, Block{FileID: -1, IsSpace: true})
			}
		}
	}
	return blocks
}

// findLeftmostSpace finds the leftmost free space in the blocks
func findLeftmostSpace(blocks []Block) int {
	for i, block := range blocks {
		if block.IsSpace {
			return i
		}
	}
	return -1
}

// findRightmostFile finds the rightmost file block and its ID
func findRightmostFile(blocks []Block) (int, int) {
	for i := len(blocks) - 1; i >= 0; i-- {
		if !blocks[i].IsSpace {
			return i, blocks[i].FileID
		}
	}
	return -1, -1
}

// moveFile moves a single block from the rightmost file to the leftmost space
func moveFile(blocks []Block) bool {
	spacePos := findLeftmostSpace(blocks)
	if spacePos == -1 {
		return false // No more spaces to fill
	}

	filePos, fileID := findRightmostFile(blocks)
	if filePos == -1 || filePos < spacePos {
		return false // No more files to move
	}

	// Move the file block
	blocks[spacePos] = Block{FileID: fileID, IsSpace: false}
	blocks[filePos] = Block{FileID: -1, IsSpace: true}
	return true
}

// calculateChecksum calculates the final checksum after compaction
func calculateChecksum(blocks []Block) int {
	checksum := 0
	for pos, block := range blocks {
		if !block.IsSpace {
			checksum += pos * block.FileID
		}
	}
	return checksum
}

func Day9Part1() {
	lines, _ := data.ReadLines("day9.txt")
	input := strings.TrimSpace(lines[0]) // Assuming single line input

	// Parse the disk map
	lengths := parseDiskMap(input)

	// Create initial block representation
	blocks := createInitialBlocks(lengths)

	// Keep moving files until no more moves are possible
	for moveFile(blocks) {
		// Continue moving files
	}

	// Calculate and return the checksum
	fmt.Println(calculateChecksum(blocks))
}

func Day9Part2() {
	_, _ = data.ReadLines("day9.txt")
}
