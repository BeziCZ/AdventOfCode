package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	//"unicode"
	"regexp"
)
type nearbyNumber struct{
	lineNum int
	sIdx int
	eIdx int 
	counted bool
	value int
}

func readFile(filename string) []string{
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
	return lines
}

func getNumbers(lines []string) []nearbyNumber{
	extractedNumbers := make([]nearbyNumber, 0)
	for idx, line := range lines{
		re := regexp.MustCompile("[0-9]+")
		numIdxs := re.FindAllStringIndex(line, -1)
		for _, numIdx := range numIdxs{
			startIdx := numIdx[0]
			endIdx := numIdx[1]
			val, err := strconv.Atoi(line[numIdx[0]:numIdx[1]])
			if err != nil{
				log.Fatal(err)
			}
			number := nearbyNumber{
				lineNum: idx,
				sIdx: startIdx,
				eIdx: endIdx,
				value: val,
				counted: false,
			}
			extractedNumbers = append(extractedNumbers, number)
		}
	}
	return extractedNumbers
}

func getNeighbours(lineIdx int, idx int, c rune, numbers []nearbyNumber) ([]nearbyNumber, int) {
	neighbours := make([]nearbyNumber,0)
	for j, num := range numbers{
		if num.lineNum >= lineIdx-1 && num.lineNum <= lineIdx+1{
			if (num.sIdx >= idx -1 && num.sIdx <= idx + 1)||(num.eIdx-1 >= idx-1 && num.eIdx-1 <= idx+1){
				neighbours = append(neighbours, num)
				numbers[j].counted = true
			}
		}
	}
	return neighbours, len(neighbours)
}

func getNeighbourValue(lineIdx int, idx int, c rune, numbers []nearbyNumber) int{
	var localSum = 0
	for j, num := range numbers{
		if num.lineNum >= lineIdx-1 && num.lineNum <= lineIdx+1{
			if (num.sIdx >= idx -1 && num.sIdx <= idx + 1)||(num.eIdx-1 >= idx-1 && num.eIdx-1 <= idx+1){
				if !num.counted {
					localSum += num.value
					numbers[j].counted = true
				} 
			}
		}
	}
	return localSum
}
func main() {
	// Part 1
	lines := readFile("./input")
	var totalSum int = 0
	var totalRatio int = 0
	numbers := getNumbers(lines)

	for i := 1; i < len(lines)-1; i++{
		fmt.Println(lines[i])
		for idx, c := range lines[i]{
			if !(c >= '0' && c <= '9') && c != '.'{
				if c == '*'{
					neig, num := getNeighbours(i, idx, c, numbers)
					if num == 2{
						totalRatio += (neig[0].value * neig[1].value)
						totalSum += neig[0].value + neig[1].value
					} else{
						totalSum += neig[0].value
					}
				} else{
					totalSum += getNeighbourValue(i, idx, c, numbers)
				}
			}
		}
	}
	fmt.Println(totalSum)
	fmt.Println(totalRatio)
}