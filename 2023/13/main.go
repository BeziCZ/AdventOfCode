package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func readFile(filename string) [][]string{
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var currentSlice []string
	var result [][]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentSlice) > 0 {
				result = append(result, currentSlice)
				currentSlice = nil
			}
		} else {
			currentSlice = append(currentSlice, line)
		}
	}

	if len(currentSlice) > 0 {
		result = append(result, currentSlice)
	}

	return result
}

func getRowRef(block []string) int{
	numOfCols := len(block[0])
	numOfRows := len(block)

	for i := 0; i < numOfRows-1; i++ {
		difference := 0
		for j := 0; j < numOfCols; j++ {
			for k := 0;;k++{
				topRow := i - k
				bottomRow := i + k + 1
				if topRow < 0 || bottomRow >= numOfRows{
					break
				}
				if block[topRow][j] != block[bottomRow][j]{
					difference++
				}
			}
		}
		if difference == 0{
			return i + 1
		}
	}
	return 0
}

func getColRef(block []string) int{
	numOfCols := len(block[0])
	numOfRows := len(block)

	for i := 0; i < numOfCols-1; i++ {
		difference := 0
		for j := 0; j < numOfRows; j++ {
			for k := 0;;k++{
				leftCol := i - k
				rightCol := i + k + 1
				if leftCol < 0 || rightCol >= numOfCols{
					break
				}
				if block[j][leftCol] != block[j][rightCol]{
					difference++
				}
			}
		}
		if difference == 0{
			return i + 1
		}
	}
	return 0
}

func getColRef2(block []string) int{
	numOfCols := len(block[0])
	numOfRows := len(block)

	for i := 0; i < numOfCols-1; i++ {
		difference := 0
		for j := 0; j < numOfRows; j++ {
			for k := 0;;k++{
				leftCol := i - k
				rightCol := i + k + 1
				if leftCol < 0 || rightCol >= numOfCols{
					break
				}
				if block[j][leftCol] != block[j][rightCol]{
					difference++
				}
			}
		}
		if difference == 1{
			return i + 1
		}
	}
	return 0
}
func getRowRef2(block []string) int{
	numOfCols := len(block[0])
	numOfRows := len(block)

	for i := 0; i < numOfRows-1; i++ {
		difference := 0
		for j := 0; j < numOfCols; j++ {
			for k := 0;;k++{
				topRow := i - k
				bottomRow := i + k + 1
				if topRow < 0 || bottomRow >= numOfRows{
					break
				}
				if block[topRow][j] != block[bottomRow][j]{
					difference++
				}
			}
		}
		if difference == 1{
			return i + 1
		}
	}
	return 0
}
func main() {
	blocks := readFile("input")
	total1 := 0
	total2 := 0
	for _, block := range blocks {
		total1 += 100*getRowRef(block) + getColRef(block)
		total2 += 100*getRowRef2(block) + getColRef2(block)
	}
	fmt.Println(total1)
	fmt.Println(total2)
}