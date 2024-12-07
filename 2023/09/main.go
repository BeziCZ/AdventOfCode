package main 

import (
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"slices"
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

func getNextValue(line []int) int{
	lastVals := make([]int, 0)
	currVals := line
	allZero := true
	nextVals := make([]int, 0)
	for {
		for i:=1; i<len(currVals); i++{
			nVal := currVals[i]-currVals[i-1]
			if nVal != 0 {
				allZero = false
			}
			nextVals = append(nextVals, nVal)
		}
		lastVals = append(lastVals, currVals[len(currVals)-1])
		currVals = nextVals
		nextVals = nil
		if allZero{
			break
		} else{
			allZero = true
		}
	}
	retVal := 0
	for i := len(lastVals)-1; i >= 0; i--{
		retVal += lastVals[i]
	}
	return retVal
}

func getPrevValue(line []int) int{
	firstVals := make([]int, 0)
	currVals := line
	allZero := true
	prevVals := make([]int, 0)
	for {
		for i:=len(currVals)-1; i>0; i--{
			pVal := currVals[i]-currVals[i-1]
			if pVal != 0 {
				allZero = false
			}
			prevVals = append(prevVals, pVal)
		}
		slices.Reverse(prevVals)
		firstVals = append(firstVals, currVals[0])
		currVals = prevVals
		prevVals = nil
		if allZero{
			break
		} else{
			allZero = true
		}
	}
	retVal := 0
	for i := len(firstVals)-1; i >= 0; i--{
		retVal = firstVals[i] - retVal
	}
	return retVal
}

func main() {
	lines := readFile("input")
	nextValueSum := 0
	prevValueSum := 0
	for _, line := range lines{
		nums := make([]int, 0)
		split := strings.Split(line," ")
		for _, v := range split {
			num, _ := strconv.Atoi(v)
			nums = append(nums,num)
		}
		nextValueSum += getNextValue(nums)
		prevValueSum += getPrevValue(nums)
	}

	fmt.Println(nextValueSum)
	fmt.Println(prevValueSum)
}