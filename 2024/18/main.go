package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Coord struct {
	x, y int
}

func bfs(grid [][]string, start Coord, end Coord) int {
	// Check if start or end is blocked
	if grid[start.x][start.y] == "#" || grid[end.x][end.y] == "#" {
		return -1
	}

	// Create queue for BFS
	queue := []struct {
		pos   Coord
		steps int
	}{{start, 0}}

	// Create visited map
	visited := make(map[Coord]bool)
	visited[start] = true

	// Directions: right, down, left, up
	directions := []Coord{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	// BFS
	for len(queue) > 0 {
		// Get current position and steps
		current := queue[0]
		queue = queue[1:]

		// Check if we reached the end
		if current.pos == end {
			return current.steps
		}

		// Try all directions
		for _, dir := range directions {
			newX := current.pos.x + dir.x
			newY := current.pos.y + dir.y
			newPos := Coord{newX, newY}

			// Check if new position is valid
			if newX >= 0 && newX < len(grid) &&
				newY >= 0 && newY < len(grid[0]) &&
				grid[newX][newY] == "." &&
				!visited[newPos] {

				queue = append(queue, struct {
					pos   Coord
					steps int
				}{newPos, current.steps + 1})
				visited[newPos] = true
			}
		}
	}

	return -1 // No path found
}

func readInput(path string) []Coord {
	var lines []Coord
	file, _ := os.Open(path)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		lines = append(lines, Coord{x, y})
	}
	return lines
}

func main() {
	bytes := readInput("input.txt")
	grid := make([][]string, 71) // Test grid 7x7 full grid 71x71
	for i := range grid {
		grid[i] = make([]string, 71)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	startClock := time.Now()
	for i := 0; i < 1024; i++ {
		currByte := bytes[i]
		grid[currByte.y][currByte.x] = "#"
	}

	start := Coord{0, 0}
	end := Coord{70, 70}
	steps := bfs(grid, start, end)
	stopClock := time.Since(startClock)
	fmt.Println("Part 1:", steps, "Time:", stopClock)

	// Part 2
	breakCoord := Coord{-1, -1}
	for i := 1024; i < len(bytes); i++ {
		currByte := bytes[i]
		grid[currByte.y][currByte.x] = "#"
		if bfs(grid, start, end) == -1 {
			breakCoord = currByte
			break
		}
	}
	fmt.Println("Part 2:", breakCoord.x, breakCoord.y, "Time:", time.Since(startClock))
}
