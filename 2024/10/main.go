package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	row, col int
}

func readInput(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var topMap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, v := range strings.Split(line, "") {
			num, _ := strconv.Atoi(v)
			row = append(row, num)
		}
		topMap = append(topMap, row)
	}
	return topMap
}

func findPaths(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	// Map to store which starting zeros can reach which endpoints
	reachableFrom := make(map[Point]map[Point]bool)

	var starts []Point
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				starts = append(starts, Point{i, j})
			}
		}
	}

	// DFS function to explore paths
	var dfs func(curr Point, start Point, prevVal int, visited map[Point]bool)
	dfs = func(curr Point, start Point, prevVal int, visited map[Point]bool) {
		if curr.row < 0 || curr.row >= rows || curr.col < 0 || curr.col >= cols {
			return
		}

		currVal := grid[curr.row][curr.col]
		if currVal != prevVal+1 {
			return
		}

		if visited[curr] {
			return
		}

		if currVal == 9 {
			// Initialize map for this endpoint if it doesn't exist
			if reachableFrom[curr] == nil {
				reachableFrom[curr] = make(map[Point]bool)
			}
			// Mark that this endpoint is reachable from the starting zero
			reachableFrom[curr][start] = true
			return
		}

		visited[curr] = true

		directions := []Point{
			{-1, 0}, // up
			{1, 0},  // down
			{0, -1}, // left
			{0, 1},  // right
		}

		for _, dir := range directions {
			next := Point{curr.row + dir.row, curr.col + dir.col}
			dfs(next, start, currVal, visited)
		}

		delete(visited, curr)
	}

	// Process each starting zero
	for _, start := range starts {
		visited := make(map[Point]bool)
		dfs(start, start, -1, visited)
	}

	totalPaths := 0
	for _, startingPoints := range reachableFrom {
		totalPaths += len(startingPoints)
	}

	return totalPaths
}

func findDistinctPaths(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	totalPaths := 0

	var starts []Point
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 0 {
				starts = append(starts, Point{i, j})
			}
		}
	}

	// DFS function to explore paths
	var dfs func(curr Point, prevVal int, visited map[Point]bool)
	dfs = func(curr Point, prevVal int, visited map[Point]bool) {
		// Base cases
		if curr.row < 0 || curr.row >= rows || curr.col < 0 || curr.col >= cols {
			return
		}

		currVal := grid[curr.row][curr.col]
		// Must increase by exactly 1
		if currVal != prevVal+1 {
			return
		}

		// If we've already visited this cell in current path
		if visited[curr] {
			return
		}

		// If we reached 9, we found a valid path
		if currVal == 9 {
			totalPaths++
			return
		}

		// Mark current cell as visited
		visited[curr] = true

		// Try all four directions
		directions := []Point{
			{-1, 0}, // up
			{1, 0},  // down
			{0, -1}, // left
			{0, 1},  // right
		}

		for _, dir := range directions {
			next := Point{curr.row + dir.row, curr.col + dir.col}
			dfs(next, currVal, visited)
		}

		// Backtrack
		delete(visited, curr)
	}

	// Start DFS from each zero
	for _, start := range starts {
		visited := make(map[Point]bool)
		dfs(start, -1, visited)
	}

	return totalPaths
}

func main() {
	topMap := readInput("input.txt")

	start := time.Now()
	paths := findPaths(topMap)
	elapsed := time.Since(start)
	fmt.Println("Part 1:", paths, "in time:", elapsed)

	start = time.Now()
	paths2 := findDistinctPaths(topMap)
	elapsed = time.Since(start)
	fmt.Println("Part 2:", paths2, "in time:", elapsed)
}
