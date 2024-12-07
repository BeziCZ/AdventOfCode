package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(path string) map[int][]int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		res, _ := strconv.Atoi(strings.Split(line, ":")[0])
		numString := strings.Split(strings.Split(line, ":")[1], " ")
		var nums []int
		for i := 1; i < len(numString); i++ {
			num, _ := strconv.Atoi(numString[i])
			nums = append(nums, num)
		}
		input[res] = nums
	}
	return input
}

func concatenate(x, y int) int {
	yString := fmt.Sprintf("%d", y)
	multiplier := 1
	for i := 0; i < len(yString); i++ {
		multiplier *= 10
	}
	return x*multiplier + y
}

func evaluate(numbers []int, operators []string) int {
	result := numbers[0]
	for i := 0; i < len(operators); i++ {
		switch operators[i] {
		case "+":
			result += numbers[i+1]
		case "*":
			result *= numbers[i+1]
		case "||":
			result = concatenate(result, numbers[i+1])
		}
	}
	return result
}

func IsPossible(numbers []int, target int, ops []string) bool {
	n := len(numbers) - 1 // number of operators needed
	operators := make([]string, n)

	// Helper function to try all possible operator combinations
	var tryAllCombinations func(pos int) bool
	tryAllCombinations = func(pos int) bool {
		// If we've filled all operator positions, evaluate
		if pos == n {
			result := evaluate(numbers, operators)
			return result == target
		}

		// Try both operators at current position
		for _, op := range ops {
			operators[pos] = op
			if tryAllCombinations(pos + 1) {
				return true
			}
		}
		return false
	}

	return tryAllCombinations(0)
}

func main() {
	input := readInput("input.txt")

	ops1 := []string{"+", "*"}
	ops2 := []string{"+", "*", "||"}
	total := 0
	start := time.Now()
	for res, nums := range input {
		if IsPossible(nums, res, ops1) {
			total += res
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1 total:", total, "in:", elapsed)

	total = 0
	start = time.Now()
	for res, nums := range input {
		if IsPossible(nums, res, ops2) {
			total += res
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Part 2 total:", total, "in:", elapsed)
}
