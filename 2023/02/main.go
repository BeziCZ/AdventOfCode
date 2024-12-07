package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"io"
)
func isPossible(s string) bool{
	nums := [3]int{12,13,14} 
	colors := [3]string{"red", "green", "blue"}
	draws := strings.Split(s,";")
	for i := 0; i < len(draws); i++ {
		cols := strings.Split(draws[i], ",")
		for j := 0; j < len(cols); j++ {
			for idx, color := range colors{
				if strings.Contains(cols[j], color){
					curr := strings.Split(cols[j], " ")
					i, _ := strconv.Atoi(curr[1])
					if nums[idx] < i{
						return false
					}
				}
			}
		}
	}
	return true
}

func gamePower(s string) int{
	mins := [3]int{0,0,0}
	colors := [3]string{"red", "green", "blue"}
	draws := strings.Split(s,";")
	for i := 0; i < len(draws); i++ {
		cols := strings.Split(draws[i], ",")
		for j := 0; j < len(cols); j++ {
			for idx, color := range colors{
				if strings.Contains(cols[j], color){
					curr := strings.Split(cols[j], " ")
					i, _ := strconv.Atoi(curr[1])
					if i > mins[idx]{
						mins[idx] = i
					}
				}
			}
		}
	}
	return (mins[0]*mins[1]*mins[2])
}

func main() {
	fmt.Println("Opening file...")
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//var sumId int = 0
	var totalPower int = 0
	for {
		line, err := reader.ReadString('\n')
		if err == nil {
			lineSplit := strings.Split(line,":")
			//gameId, _ := strconv.Atoi(strings.Split(lineSplit[0]," ")[1])
			//if isPossible(lineSplit[1]){
			//	sumId += gameId
			//}
			totalPower += gamePower(lineSplit[1])
		}
		if err != nil{
			if err == io.EOF {
				//lineSplit := strings.Split(line,":")
				//gameId, _ := strconv.Atoi(strings.Split(lineSplit[0]," ")[1])
				//if isPossible(lineSplit[1]){
				//	sumId += gameId
				//}
				//totalPower += gamePower(lineSplit[1])
				break
			}
			log.Fatal(err)
			return
		}
	}
	//fmt.Println(sumId)
	fmt.Println(totalPower)
}
