package main

import (
	"fmt"
	"os"
	"strings"
	"log"
	"slices"
	"regexp"
	"io/ioutil"
	"strconv"
)

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
func getCardPoints(line string) int{
	numbers := strings.Split(line, "|")
	re := regexp.MustCompile("[0-9]+")
	winning := re.FindAllString(numbers[0],-1)
	random := re.FindAllString(numbers[1],-1)
	var worth int = 0 
	for _, win := range random{
		if slices.Contains(winning,win){ 
			if worth == 0{
				worth = 1
			} else {
				worth *= 2
			}
		}
	}
	return worth
}

func getNumberOfNextCards(line string) (int, int){
	game := strings.Split(line, ":")
	numbers := strings.Split(game[1], "|")
	re := regexp.MustCompile("[0-9]+")
	id, _ := strconv.Atoi(re.FindAllString(game[0],-1)[0])
	winning := re.FindAllString(numbers[0],-1)
	random := re.FindAllString(numbers[1],-1)
	var numOfNext int = 0 
	for _, win := range random{
		if slices.Contains(winning,win){
			numOfNext++
		}
	}
	return id, numOfNext
}
func main() {
	lines := readFile("./input")
	//Part 1
	var totalValue int = 0
	for _, line := range lines{
		if line != "" {
		split_line := strings.Split(line,":")
		totalValue += getCardPoints(split_line[1])
		}
	}
	fmt.Println(totalValue) 
	// Part 2
	numsOfCards := map[int]int{1:1}
	for _, line := range lines{
		if line != ""{
			id, numOfNext := getNumberOfNextCards(line)
			if  _, good := numsOfCards[id]; !good{
				numsOfCards[id] = 1
			}
			if numOfNext == 0{
				continue
			}
			for i := id+1; i <= id+numOfNext; i++ {
				if  _, good := numsOfCards[i]; !good{
					numsOfCards[i] = 1
				}
				numsOfCards[i] += numsOfCards[id]
			}
		}
	}
	var totalNumberOfCards int = 0
	for _, n := range numsOfCards{
		totalNumberOfCards += n
	}
	fmt.Println(totalNumberOfCards)
}