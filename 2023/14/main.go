package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"time"
)

var data = make([][]string,0)
const rock = "O"
const empty = "."

func readFile(filename string) [][]string{
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(raw), "\n")
	matrix := make([][]string, 0)
	for _, line := range lines {
		matrix = append(matrix, strings.Split(line,""))
	}
	return matrix
}

func rollNorth(height, width int){
	for col := 0; col < width; col++{
		isFirst := true
		lastRock := -1
		for row := 1; row < height; row++{
			if isFirst && data[row][col] == rock{
				currRow := row
				for currRow - 1 >= 0 && data[currRow-1][col] == empty{
					currRow--
				}
				if currRow != row{
					data[currRow][col] = rock
					data[row][col] = empty
				}
				lastRock = currRow
				isFirst = false
			} else if data[row][col] == rock{
				if data[lastRock+1][col] == empty{
					data[lastRock+1][col] = rock
					data[row][col] = empty
				}
				lastRock++
			} else if data[row][col] != rock && data[row][col] != empty {
				isFirst = true
			}
		}
	}
}

func rollSouth(height, width int){
	for col := 0; col < width; col++{
		lastRock := -1
		isFirst := true
		for row := height-1; row >= 0; row--{
			if isFirst && data[row][col] == rock{
				currRow := row
				for currRow + 1 < height && data[currRow+1][col] == empty{
					currRow++
				}
				if currRow != row{
					data[currRow][col] = rock
					data[row][col] = empty
				}
				lastRock = currRow
				isFirst = false
			} else if data[row][col] == rock{
				if data[lastRock-1][col] == empty{
					data[lastRock-1][col] = rock
					data[row][col] = empty
				}
				lastRock--
			} else if data[row][col] != rock && data[row][col] != empty {
				isFirst = true
			}
		}
	}
}

func rollWest(height, width int){
	for row := 0; row < height; row++{
		isFirst := true
		lastRock := -1
		for col := 1; col < width; col++{
			if isFirst && data[row][col] == rock{
				currCol := col
				for currCol - 1 >= 0 && data[row][currCol-1] == empty{
					currCol--
				}
				if currCol != col{
					data[row][currCol] = rock
					data[row][col] = empty
				}
				lastRock = currCol
				isFirst = false
			} else if data[row][col] == rock{
				if data[row][lastRock+1] == empty{
					data[row][lastRock+1] = rock
					data[row][col] = empty
				}
				lastRock++
			} else if data[row][col] != rock && data[row][col] != empty {
				isFirst = true
			}
		}
	}
}

func rollEast(height, width int){
	for row := 0; row < height; row++{
		lastRock := -1
		isFirst := true
		for col := width-1; col >= 0; col--{
			if isFirst && data[row][col] == rock{
				currCol := col
				for currCol + 1 < width && data[row][currCol+1] == empty{
					currCol++
				}
				if currCol != col{
					data[row][currCol] = rock
					data[row][col] = empty
				}
				lastRock = currCol
				isFirst = false
			} else if data[row][col] == rock{
				if data[row][lastRock-1] == empty{
					data[row][lastRock-1] = rock
					data[row][col] = empty
				}
				lastRock--
			} else if data[row][col] != rock && data[row][col] != empty {
				isFirst = true
			}
		}
	}
}

func getString(f [][]string) string{
	temp := make([]string,0)
	res := ""
	for _, row := range f{
		temp = append(temp,strings.Join(row,""))
	}
	for _, s := range temp{
		res += s
	}
	return res
}

func main() {
	data = readFile("testinput")
	height := len(data)
	width := len(data[0])

	stateMap := make(map[string]int,0)
	idxMap := make(map[int][][]string,0)
	s, m := 0,0

	start := time.Now()
	for i := 0; i < 1000000000; i++{
		fmt.Println(i)
		for _, line := range data{
			fmt.Println(line)
		}
		rollNorth(height, width)
		rollWest(height, width)
		rollSouth(height, width)
		rollEast(height, width)
		stringData := getString(data)
		if _, ok := stateMap[stringData]; ok {
			s = stateMap[stringData]
			m = i - s
			break
		}
		stateMap[stringData] = i
		idxMap[i] = data
	}
	fmt.Println(time.Since(start))
	finalState := idxMap[s + ((1000000000 - s) % m)-1]
	fmt.Println(finalState[0][0])
	for _, line := range idxMap[0]{
		fmt.Println(line)
	}
	total := 0
	for i := 0; i < height; i++{
		for j := 0; j < width; j++{
			if idxMap[0][i][j] == "O"{
				fmt.Println(total)
				total += height - i
			}
		}
	}
	fmt.Println(total)
}