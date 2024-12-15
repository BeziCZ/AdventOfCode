package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type antenna struct {
	xCoord int
	yCoord int
}

func readInput(path string) [][]rune {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		var runes []rune
		for _, s := range split {
			runes = append(runes, []rune(s)...)
		}
		input = append(input, runes)
	}
	return input
}

func findThemBitches(input [][]rune) map[rune][]antenna {
	var antennas map[rune][]antenna = make(map[rune][]antenna)

	for i, row := range input {
		for j, cell := range row {
			if cell != '.' && cell != '#' {
				if antennas[cell] == nil {
					antennas[cell] = []antenna{{xCoord: i, yCoord: j}}
				} else {
					antennas[cell] = append(antennas[cell], antenna{xCoord: i, yCoord: j})
				}
			}
		}
	}
	return antennas
}

func main() {
	path := "input.txt"
	input := readInput(path)

	//Part 1
	bitches := findThemBitches(input)
	overlap := 0
	for _, coords := range bitches {
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				diffX := coords[j].xCoord - coords[i].xCoord
				diffY := coords[j].yCoord - coords[i].yCoord
				if coords[i].xCoord-diffX < len(input) &&
					coords[i].xCoord-diffX >= 0 &&
					coords[i].yCoord-diffY < len(input[0]) &&
					coords[i].yCoord-diffY >= 0 {
					if input[coords[i].xCoord-diffX][coords[i].yCoord-diffY] == '.' {
						input[coords[i].xCoord-diffX][coords[i].yCoord-diffY] = '#'
					} else if input[coords[i].xCoord-diffX][coords[i].yCoord-diffY] == '#' {
						overlap++
					}
				}
				if coords[j].xCoord+diffX < len(input) &&
					coords[j].xCoord+diffX >= 0 &&
					coords[j].yCoord+diffY < len(input[0]) &&
					coords[j].yCoord+diffY >= 0 {
					if input[coords[j].xCoord+diffX][coords[j].yCoord+diffY] == '.' {
						input[coords[j].xCoord+diffX][coords[j].yCoord+diffY] = '#'
					} else if input[coords[j].xCoord+diffX][coords[j].yCoord+diffY] == '#' {
						overlap++
					}
				}
			}
		}
	}

	total := 0
	for _, row := range input {
		fmt.Println(string(row))
		for _, cell := range row {
			if cell == '#' {
				total++
			}
		}
	}
	fmt.Println(total + overlap - 1)

	//Part 2
	input = readInput(path)
	total = 0
	for _, coords := range bitches {
		for i := 0; i < len(coords)-1; i++ {
			for j := i + 1; j < len(coords); j++ {
				diffX := coords[j].xCoord - coords[i].xCoord
				diffY := coords[j].yCoord - coords[i].yCoord
				k := 0
				l := 0
				for {
					if coords[i].xCoord-k*diffX < len(input) &&
						coords[i].xCoord-k*diffX >= 0 &&
						coords[i].yCoord-k*diffY < len(input[0]) &&
						coords[i].yCoord-k*diffY >= 0 {
						if input[coords[i].xCoord-k*diffX][coords[i].yCoord-k*diffY] == '.' {
							input[coords[i].xCoord-k*diffX][coords[i].yCoord-k*diffY] = '#'
						} else if input[coords[i].xCoord-k*diffX][coords[i].yCoord-k*diffY] == '#' {
							total++
						}
					} else {
						break
					}
					k++
				}
				for {
					if coords[j].xCoord+l*diffX < len(input) &&
						coords[j].xCoord+l*diffX >= 0 &&
						coords[j].yCoord+l*diffY < len(input[0]) &&
						coords[j].yCoord+l*diffY >= 0 {
						if input[coords[j].xCoord+l*diffX][coords[j].yCoord+l*diffY] == '.' {
							input[coords[j].xCoord+l*diffX][coords[j].yCoord+l*diffY] = '#'
						} else if input[coords[j].xCoord+l*diffX][coords[j].yCoord+l*diffY] == '#' {
							total++
						}
					} else {
						break
					}
					l++
				}
			}
		}
	}

	for _, row := range input {
		fmt.Println(string(row))
		for _, cell := range row {
			if cell == '#' {
				total++
			}
		}
	}
	fmt.Println(total)
}
