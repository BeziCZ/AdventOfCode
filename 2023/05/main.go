package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"unicode"
	"slices"
)

type Interval struct {
	start int
	end int
}

var (
	seed2soil = make(map[Interval]Interval,0)
	soil2fert = make(map[Interval]Interval,0)
	fert2wat = make(map[Interval]Interval,0)
	wat2ligh = make(map[Interval]Interval,0)
	ligh2temp = make(map[Interval]Interval,0)
	temp2hum = make(map[Interval]Interval,0)
	hum2loc = make(map[Interval]Interval,0)
)
func readFile(filename string) []string{
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func getWantedSeeds1(line string) []int{
	seeds := make([]int,0)
	seedNums := strings.Split(strings.Trim(strings.Split(line, ":")[1]," "), " ")
	for _, seed := range seedNums {
		num, _ := strconv.Atoi(seed)
		seeds = append(seeds, num)
	}
	return seeds
}

func getWantedSeeds2(line string) []Interval {
	seeds := make([]Interval,0)
	seedNums := strings.Split(strings.Trim(strings.Split(line, ":")[1]," "), " ")
	tmp := Interval{start: 0, end: 0}
	for i, num := range seedNums{
		if i % 2 != 0{
			n, _ := strconv.Atoi(num)
			tmp.end = tmp.start + n -1
			seeds = append(seeds, tmp)
			tmp.start = 0
			tmp.end = 0
		} else{
			n, _ := strconv.Atoi(num)
			tmp.start = n
		}
	}
	return seeds
}

func getNext(num int, idx int) int {
	nextNum := -1
	
	switch idx{
	case 0:
	{
		for source, dest := range seed2soil{
			if source.start <= num && num <= source.end {
				nextNum = dest.start + (num-source.start)
			}
		}
		if nextNum == -1 { return num}
	}
	case 1:
	{
		for source, dest := range soil2fert{
			if source.start <= num && num <= source.end {
				nextNum = dest.start + (num-source.start)
			}
		}
		if nextNum == -1 { return num}
	}
	case 2:
	{
		for source, dest := range fert2wat{
			if source.start <= num && num <= source.end {
				nextNum = dest.start + (num-source.start)
			}
		}
		if nextNum == -1 { return num}
	} 
	case 3:
	{
		for source, dest := range wat2ligh{
			if source.start <= num && num <= source.end {
				nextNum = dest.start + (num-source.start)
			}
		}
		if nextNum == -1 { return num }
	} 
	case 4:
	{
		for source, dest := range ligh2temp{
			if source.start <= num && num <= source.end {
				nextNum = dest.start + (num-source.start)
			}
		}
		if nextNum == -1 { return num }
	} 
	case 5:
	{
		for source, dest := range temp2hum{
			if source.start <= num && num <= source.end {
				nextNum = dest.start + (num-source.start)
			}
		}
		if nextNum == -1 { return num }
	} 
	case 6:
	{
		for source, dest := range hum2loc{
			if source.start <= num && num <= source.end {
				nextNum = dest.start + (num-source.start)
			}
		}
		if nextNum == -1 { return num }
	}
	}
	return nextNum
}

func main() {
	lines := readFile("input")
	wantedSeeds := getWantedSeeds1(lines[0])
	wantedSeeds2 := getWantedSeeds2(lines[0])
	lastLoading := 0
	for i := 2; i < len(lines); i++ {
		if lines[i] == ""{
			continue
		} else if unicode.IsDigit(rune(lines[i][0])){
			split := strings.Split(lines[i]," ")
			d, _ := strconv.Atoi(string(split[0]))
			s, _ := strconv.Atoi(string(split[1]))
			r, _ := strconv.Atoi(string(split[2]))
			source := Interval{start: s, end: s+r-1}
			dest := Interval{start: d, end: d+r-1}
			switch lastLoading{
			case 1:
				seed2soil[source] = dest
			case 2:
				soil2fert[source] = dest
			case 3:
				fert2wat[source] = dest
			case 4:
				wat2ligh[source] = dest
			case 5:
				ligh2temp[source] = dest
			case 6:
				temp2hum[source] = dest
			case 7:
				hum2loc[source] = dest
			}
		} else{
			lastLoading++
		}
	}
	locations := make([]int,0)
	//Part1
	for _, s := range wantedSeeds {
		currNum := s
		for i := 0; i < 7; i++ {
			currNum = getNext(currNum, i)
		}
		locations = append(locations, currNum)
	}
	minLoc := slices.Min(locations)
	fmt.Println(minLoc)
	//Part2
	location2 := 0
	for idx, s := range wantedSeeds2 {
		fmt.Println(idx)
		for i := s.start; i <= s.end; i++ {
			currNum := i
			for j := 0; j < 7; j++ {
				currNum = getNext(currNum, j)
			}
			if idx == 0 && i == 0{
				location2 = currNum
			} else{
				if location2 > currNum {
					location2 = currNum
				}
			}
		}
	}
	fmt.Println(location2)
}