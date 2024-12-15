package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	lines := (strings.Split(string(file), "\n"))
	return strings.Split(lines[0], "")
}

func createSlice1(line []string) []string {
	id := 0
	gappedSlice := make([]string, 0)
	for i := range line {
		if i%2 == 0 {
			id = i / 2
			fileSpace, _ := strconv.ParseInt(line[i], 10, 32)
			for j := 0; j < int(fileSpace); j++ {
				gappedSlice = append(gappedSlice, fmt.Sprintf("%d", id))
			}
		} else {
			freeSpace, _ := strconv.ParseInt(line[i], 10, 32)
			for j := 0; j < int(freeSpace); j++ {
				gappedSlice = append(gappedSlice, ".")
			}
		}
	}
	return gappedSlice
}

func part1move(gappedSlice []string) []string {
	for i, j := 0, len(gappedSlice)-1; i < j; {
		if gappedSlice[i] == "." && gappedSlice[j] != "." {
			gappedSlice[i], gappedSlice[j] = gappedSlice[j], gappedSlice[i]
			i++
			j--
		} else {
			if gappedSlice[i] != "." {
				i++
			} else if gappedSlice[j] != "." {
				j--
			} else if gappedSlice[i] == "." && gappedSlice[j] == "." {
				j--
			}
		}
	}
	return gappedSlice
}

func calculateChecksum(disk []string) int64 {
	sum := int64(0)
	for i, r := range disk {
		if r != "." {
			value, _ := strconv.ParseInt(disk[i], 10, 32)
			sum += int64(i) * value
		}
	}
	return sum
}

func main() {
	numbers := readInput("input.txt")

	start := time.Now()
	gappedSlice := createSlice1(numbers)
	part1slice := part1move(gappedSlice)

	checksum := calculateChecksum(part1slice)
	elapsed := time.Since(start)
	fmt.Println("Part 1:", checksum, "in:", elapsed)

}
