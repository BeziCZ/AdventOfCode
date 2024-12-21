package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var cache = map[string]int{}

func ways(combination string, patterns []string) (n int) {
	if val, ok := cache[combination]; ok {
		return val
	}

	defer func() {
		cache[combination] = n
	}()

	if combination == "" {
		return 1
	}

	for _, pattern := range patterns {
		if strings.HasPrefix(combination, pattern) {
			n += ways(combination[len(pattern):], patterns)
		}
	}
	return n
}

func main() {
	input, _ := os.ReadFile("testinput.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	towels := strings.Split(split[0], ", ")
	start := time.Now()
	count := 0

	for _, combination := range strings.Fields(split[1]) {
		if n := ways(combination, towels); n > 0 {
			count++
		}
	}
	fmt.Println("Part 1:", count, time.Since(start))

	start = time.Now()
	count = 0
	for _, combination := range strings.Fields(split[1]) {
		if n := ways(combination, towels); n > 0 {
			count += n
		}
	}
	fmt.Println("Part 2:", count, time.Since(start))
}
