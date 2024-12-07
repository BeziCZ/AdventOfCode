package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"io/ioutil"
	"slices"
	"math"
)

var (
	starMap = make([][]string,0)
	galaxyCoord = make(map[int][]int,0) // idx : [X,Y]
	emptyCols = make([]int,0)
	emptyRows = make([]int,0)
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

func getNumberOfEmptyRowsOrCols(coords []int, rc int) (int){
	numOfEmpty := 0
	switch rc{	
	case 0:
	{
		for _, row := range emptyRows{
			if slices.Min(coords) < row &&  slices.Max(coords) > row{
				numOfEmpty++
			}
		}
	}
	case 1:
	{
		for _, col := range emptyCols{
			if slices.Min(coords) < col &&  slices.Max(coords) > col{
				numOfEmpty++
			}
		}	
	}
	}
	return numOfEmpty
}
func main() {
	lines := readFile("input")
	idx := 0
	for i, line := range lines{
		if strings.Contains(line,"#"){
			starMap = append(starMap, strings.Split(line, ""))
		} else {
			starMap = append(starMap, strings.Split(line,""))
			emptyRows = append(emptyRows,i)
		}
	}
	for i := 0; i < len(starMap[0]); i++{
		isEmpty := true
		for j := 0; j < len(starMap); j++ {
			if starMap[j][i] == "#"{
				isEmpty = false
			}
		}
		if isEmpty{
			emptyCols = append(emptyCols,i)
		}
	}

	for i, local := range starMap{
		for j, gal := range local{
			if gal == "#"{
				galaxyCoord[idx] = []int{j,i}
				idx++
			}
		}
	}

	totalDist := 0

	for i:= 0; i < len(galaxyCoord); i++ {
		for j:= i+1; j < len(galaxyCoord); j++{
			if i != j{
				numOfEmptyRows := getNumberOfEmptyRowsOrCols([]int{galaxyCoord[i][1],galaxyCoord[j][1]},0)
				numOfEmptyCols := getNumberOfEmptyRowsOrCols([]int{galaxyCoord[i][0],galaxyCoord[j][0]},1)
				normalDist := int((math.Abs(float64(galaxyCoord[j][0]-galaxyCoord[i][0])))+(math.Abs(float64(galaxyCoord[j][1]-galaxyCoord[i][1]))))
				totalDist += (normalDist + numOfEmptyCols*1000000 + numOfEmptyRows*1000000)-numOfEmptyCols-numOfEmptyRows
			}
		}
	}
	fmt.Println(totalDist)
}