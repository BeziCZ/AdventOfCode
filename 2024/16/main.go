package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

func readInput(path string) [][]string {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, strings.Split(line, ""))
	}
	return data
}

type Point struct {
	x, y int
}

type State struct {
	pos       Point
	direction int // 0: right, 1: down, 2: left, 3: up
	cost      int
}

type PQItem struct {
	state    State
	priority int
	index    int
}

type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PQItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func findShortestPath(grid [][]string) int {
	// Find start and end points
	var start, end Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "S" {
				start = Point{x, y}
			} else if grid[y][x] == "E" {
				end = Point{x, y}
			}
		}
	}

	// Directions: right, down, left, up
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	// Initialize priority queue with starting state
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &PQItem{
		state:    State{start, 0, 0}, // Start facing right
		priority: 0,
	})

	// Keep track of visited states
	visited := make(map[string]bool)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*PQItem)
		curr := item.state

		// Generate unique key for current state
		key := fmt.Sprintf("%d,%d,%d", curr.pos.x, curr.pos.y, curr.direction)
		if visited[key] {
			continue
		}
		visited[key] = true

		// Check if we reached the end
		if curr.pos == end {
			return curr.cost
		}

		// Try all possible directions
		for newDir := 0; newDir < 4; newDir++ {
			newX := curr.pos.x + dx[newDir]
			newY := curr.pos.y + dy[newDir]

			// Check if the new position is valid
			if newX < 0 || newY < 0 || newY >= len(grid) || newX >= len(grid[newY]) || grid[newY][newX] == "#" {
				continue
			}

			// Calculate new cost
			turnCost := 0
			if curr.direction != newDir {
				turnCost = 1000 // Cost for turning
			}
			newCost := curr.cost + 1 + turnCost // 1 for step + turn cost if any

			heap.Push(&pq, &PQItem{
				state:    State{Point{newX, newY}, newDir, newCost},
				priority: newCost,
			})
		}
	}

	return math.MaxInt32 // No path found
}

func main() {
	data := readInput("input.txt")

	//startClock := time.Now()
	result := findShortestPath(data)
	if result == math.MaxInt32 {
		fmt.Println("No path found")
	} else {
		fmt.Println("Shortest path cost:", result)
	}
}
