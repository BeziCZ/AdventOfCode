package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func mix(secretNumber, value int) int {
	return secretNumber ^ value
}

func prune(secretNumber int) int {
	return secretNumber % 16777216
}

func transformSecretNumber(secretNumber int) int {
	// Step 1: Multiply by 64, mix, and prune
	step1Result := secretNumber * 64
	secretNumber = mix(secretNumber, step1Result)
	secretNumber = prune(secretNumber)

	// Step 2: Divide by 32 (floor), mix, and prune
	step2Result := int(math.Floor(float64(secretNumber) / 32))
	secretNumber = mix(secretNumber, step2Result)
	secretNumber = prune(secretNumber)

	// Step 3: Multiply by 2048, mix, and prune
	step3Result := secretNumber * 2048
	secretNumber = mix(secretNumber, step3Result)
	secretNumber = prune(secretNumber)

	return secretNumber
}

func getTransformedNumbers(numbers []int, numOfTransforms int) (int, [][]int) {
	prices := [][]int{}
	res := 0
	for _, num := range numbers {
		prices = append(prices, []int{num % 10})
		for i := 0; i < numOfTransforms; i++ {
			num = transformSecretNumber(num)
			prices[len(prices)-1] = append(prices[len(prices)-1], num%10)
		}
		res += num
	}
	return res, prices
}

type Diff [4]int

func analyzePriceDifferences(priceLists [][]int) int {
	diffTable := make(map[Diff]int)

	for _, prices := range priceLists {
		seen := make(map[Diff]bool)

		for i := 4; i < len(prices); i++ {
			diff := Diff{
				prices[i-3] - prices[i-4],
				prices[i-2] - prices[i-3],
				prices[i-1] - prices[i-2],
				prices[i] - prices[i-1],
			}

			if seen[diff] {
				continue
			}

			diffTable[diff] += prices[i]
			seen[diff] = true
		}
	}

	maxVal := 0
	for _, val := range diffTable {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func main() {
	numbers := []int{}
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		numbers = append(numbers, num)
	}

	start := time.Now()
	sum, prices := getTransformedNumbers(numbers, 2000)
	elapsed := time.Since(start)
	fmt.Println("The sum is: ", sum, "Time:", elapsed)

	start = time.Now()
	maxBananas := analyzePriceDifferences(prices)
	elapsed = time.Since(start)
	fmt.Println("The maximum number of bananas is: ", maxBananas, "Time:", elapsed)
}
