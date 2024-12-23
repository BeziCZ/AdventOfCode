package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type Tile struct {
	val     string
	visited bool
}

type Coord struct {
	x int
	y int
}

type Edge struct {
	from, to Coord
}

func readInput(path string) [][]Tile {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]Tile
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := []Tile{}
		line := scanner.Text()
		split := strings.Split(line, "")
		for _, val := range split {
			row = append(row, Tile{val, false})
		}
		data = append(data, row)
	}
	return data
}

func findRegions(data [][]Tile) map[string][][]Coord {
	regions := make(map[string][][]Coord)
	fill := func(x, y int, val string) []Coord {
		var queue deque.Deque[Coord]
		queue.PushBack(Coord{x, y})
		region := make([]Coord, 0)
		region = append(region, Coord{x, y})

		for {
			if queue.Len() == 0 {
				break
			}
			coord := queue.PopFront()
			direction := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
			for _, dir := range direction {
				newX := coord.x + dir.x
				newY := coord.y + dir.y
				if newX >= 0 && newX < len(data) && newY >= 0 && newY < len(data[0]) {
					if data[newX][newY].val == val && !data[newX][newY].visited {
						data[newX][newY].visited = true
						queue.PushBack(Coord{newX, newY})
						region = append(region, Coord{newX, newY})
					}
				}
			}
		}
		return region
	}

	for x := range data {
		for y := range data[0] {
			if !data[x][y].visited {
				data[x][y].visited = true
				region := fill(x, y, data[x][y].val)
				regions[data[x][y].val] = append(regions[data[x][y].val], region)
			}
		}
	}

	return regions
}

func calculatePerimeter(region []Coord) int {
	perimeter := 0
	direction := []Coord{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, coord := range region {
		numOfEdges := 0
		for _, dir := range direction {
			newX := coord.x + dir.x
			newY := coord.y + dir.y
			if !slices.Contains(region, Coord{newX, newY}) {
				numOfEdges++
			}
		}
		perimeter += numOfEdges
	}
	return perimeter
}

func traceBoundaries(region []Coord) int {
	edges := 0
	for _, c := range region {
		north := Coord{c.x - 1, c.y}
		west := Coord{c.x, c.y - 1}
		northWest := Coord{c.x - 1, c.y - 1}

		if !slices.Contains(region, north) {
			sameEdge := slices.Contains(region, west) && !slices.Contains(region, northWest)
			if !sameEdge {
				edges++
			}
		}

		south := Coord{c.x + 1, c.y}
		southWest := Coord{c.x + 1, c.y - 1}

		if !slices.Contains(region, south) {
			sameEdge := slices.Contains(region, west) && !slices.Contains(region, southWest)
			if !sameEdge {
				edges++
			}
		}

		if !slices.Contains(region, west) {
			sameEdge := slices.Contains(region, north) && !slices.Contains(region, northWest)
			if !sameEdge {
				edges++
			}
		}

		east := Coord{c.x, c.y + 1}
		northEast := Coord{c.x - 1, c.y + 1}
		if !slices.Contains(region, east) {
			sameEdge := slices.Contains(region, north) && !slices.Contains(region, northEast)
			if !sameEdge {
				edges++
			}
		}
	}
	return edges
}

func part1(data [][]Tile) int {
	total := 0
	regions := findRegions(data)

	for _, regionList := range regions {
		for _, region := range regionList {
			area := len(region)
			perimeter := calculatePerimeter(region)
			total += area * perimeter
		}
	}

	return total
}

func part2(data [][]Tile) int {
	regions := findRegions(data)
	total := 0
	for _, regionList := range regions {
		for _, region := range regionList {
			area := len(region)
			sides := traceBoundaries(region)
			total += sides * area
		}
	}
	return total
}

func main() {
	data := readInput("input.txt")

	start := time.Now()
	part1 := part1(data)
	elapsed := time.Since(start)
	fmt.Println("Part 1:", part1, "Time:", elapsed)

	// Part 2
	data = readInput("input.txt")
	start = time.Now()
	part2 := part2(data)
	elapsed = time.Since(start)
	fmt.Println("Part 2:", part2, "Time:", elapsed)

}
