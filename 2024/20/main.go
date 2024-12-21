package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func readInput(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		grid = append(grid, split)
	}

	return grid
}

type Point struct {
	x, y int
}

// Direction vectors for up, right, down, left
var dx = []int{-1, 0, 1, 0}
var dy = []int{0, 1, 0, -1}

func findPath(grid [][]string) []Point {
	rows, cols := len(grid), len(grid[0])

	// Find start and end points
	var start, end Point
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == "S" {
				start = Point{i, j}
			} else if grid[i][j] == "E" {
				end = Point{i, j}
			}
		}
	}

	// Queue for BFS
	queue := []Point{start}

	// Keep track of visited cells and their previous cell
	visited := make(map[Point]bool)
	parent := make(map[Point]Point)
	visited[start] = true

	// BFS
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check if we reached the end
		if current == end {
			break
		}

		// Try all four directions
		for i := 0; i < 4; i++ {
			newX, newY := current.x+dx[i], current.y+dy[i]
			next := Point{newX, newY}

			// Check if the new position is valid and unvisited
			if newX >= 0 && newX < rows && newY >= 0 && newY < cols &&
				!visited[next] && (grid[newX][newY] == "." || grid[newX][newY] == "E") {
				queue = append(queue, next)
				visited[next] = true
				parent[next] = current
			}
		}
	}

	// Reconstruct path
	path := []Point{}
	if !visited[end] {
		return path // Return empty path if end wasn't reached
	}

	current := end
	for current != start {
		path = append([]Point{current}, path...) // Prepend to path
		current = parent[current]
	}
	path = append([]Point{start}, path...)

	return path
}

func findCheats(path []Point, cheatLength int) int {
	total := 0
	for i, s := range path[:len(path)-100] {
		for j, e := range path[i+100:] {
			diffX := int(math.Abs(float64(e.x - s.x)))
			diffY := int(math.Abs(float64(e.y - s.y)))
			manhDist := diffX + diffY
			if manhDist <= cheatLength && manhDist <= j {
				total++
			}
		}
	}
	return total
}

func main() {
	// Read input file
	maze := readInput("input.txt")

	// Find path
	start := time.Now()
	path := findPath(maze)
	elapsed := time.Since(start)
	fmt.Println("Time to find path:", elapsed)

	// Find cheats
	start = time.Now()
	total := findCheats(path, 2)
	elapsed = time.Since(start)
	fmt.Println("100ps savings with cheat length 2:", total, "time:", elapsed)

	// Find cheats 2
	start = time.Now()
	total = findCheats(path, 20)
	elapsed = time.Since(start)
	fmt.Println("100ps savings with cheat length 20:", total, "time:", elapsed)
}
