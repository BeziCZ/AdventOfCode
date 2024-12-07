package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
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

func getPossibleWins(data map[int]int) int{
	var posWins int = 1
	for time, dist := range data{
		currWins := 0
		for i := 0; i < time+1; i++{
			currDist := (time-i)*i
			if currDist > dist{
				currWins++
			}
		}
		posWins *= currWins
	}
	return posWins
}

func main(){
	lines := readFile("./input")
	re := regexp.MustCompile("[0-9]+")

	times := re.FindAllString(strings.Split(lines[0],":")[1], -1)
	distance := re.FindAllString(strings.Split(lines[1],":")[1], -1)
	//Part 1
	data := make(map[int]int)

	for i := 0; i < len(times); i++{
		t, e := strconv.Atoi(times[i])
		if e != nil{
			log.Fatal(e)
		}
		d, e := strconv.Atoi(distance[i])
		data[t] = d
	}
	possibleWins := getPossibleWins(data)
	fmt.Printf("Part 1 result: %d\n", possibleWins)
	//Part 2
	totalTime := ""
	totalDist := ""
	for _, time := range times{
		totalTime += time
	}
	for _, dist := range distance{
		totalDist += dist
	}
	time, e := strconv.Atoi(totalTime)
	if e != nil{
		log.Fatal(e)
	}
	dist, e := strconv.Atoi(totalDist)
	if e != nil{
		log.Fatal(e)
	}
	totalPosWins := 0
	for i := 0; i <= time; i++{
		currDist := (time-i)*i
			if currDist > dist{
				totalPosWins++
			}
	}
	fmt.Printf("Part 2 result: %d\n", totalPosWins)
}