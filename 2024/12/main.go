package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		data = append(data, split)
	}
	return data
}

func main() {
	data := readInput("testinput.txt")
	fmt.Println(data)
}
