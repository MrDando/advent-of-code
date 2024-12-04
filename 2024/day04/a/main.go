package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Direction struct {
	dx, dy int
}

var directions = []Direction{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // diagonal down-right
	{1, -1},  // diagonal down-left
	{0, -1},  // left
	{-1, 0},  // up
	{-1, 1},  // diagonal up-right
	{-1, -1}, // diagonal up-left
}

func main() {
	const target = "XMAS"
	grid := getData()
	count := 0
	for i, line := range grid {
		for j, cell := range line {
			if cell == rune(target[0]) {
				count += checkDirections(grid, i, j, target)
			}
		}
	}

	fmt.Printf("Total count: %d\n", count)
}

func checkDirections(grid [][]rune, x, y int, target string) int {
	matches := 0
	for _, dir := range directions {
		if !isInBounds(grid, x, y, target, dir) {
			continue
		}
		if checkWord(grid, x, y, target, dir) {
			matches++
		}
	}
	return matches
}

func isInBounds(grid [][]rune, x, y int, target string, dir Direction) bool {
	lastX := x + dir.dx*(len(target)-1)
	lastY := y + dir.dy*(len(target)-1)
	return lastX >= 0 && lastY >= 0 && lastX < len(grid) && lastY < len(grid[0])
}

func checkWord(grid [][]rune, x, y int, target string, dir Direction) bool {
	for i := 1; i < len(target); i++ {
		currX := x + dir.dx*i
		currY := y + dir.dy*i
		if grid[currX][currY] != rune(target[i]) {
			return false
		}
	}
	return true
}

func getData() [][]rune {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}
