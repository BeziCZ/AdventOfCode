package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"
)

type Vertex struct{
	id int
	leftNeigh string
	rightNeigh string
	value string
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

func GCD(a int, b int) int{
	for b != 0{
		a, b = b, a%b
	}
	return a
}

func LCM(a int, b int) int{
	return (a*b)/GCD(a,b)
}
func goTroughGraph1(inst string, graph map[string]Vertex, start string) int{
	currNode := graph[start]
	numberOfSteps := 0
	for {
		for _, c := range inst{
			if c == 'L'{
				currNode = graph[currNode.leftNeigh]
				numberOfSteps++
			} else if c == 'R'{
				currNode = graph[currNode.rightNeigh]
				numberOfSteps++
			}
			if currNode.value == "ZZZ"{
				return numberOfSteps
			}
		}
	}
}

func goTroughGraph2(inst string, graph map[string]Vertex, start []Vertex) int{
	totalSteps := make([]int,0)
	lInst := len(inst)
	for _, v := range start{
		currVert := v
		numberOfSteps := 0
		currInst := 0
		for !(currVert.value[2] == 'Z'){
			switch c := inst[currInst]; c{
			case 'L':
				currVert = graph[currVert.leftNeigh]
			case 'R':
				currVert = graph[currVert.rightNeigh]
			}
			numberOfSteps++
			currInst = (currInst+1) % lInst
		}
		totalSteps = append(totalSteps, numberOfSteps)
	}
	res := 1
	for _, step := range totalSteps{
		res = LCM(res,step)
	}
	return res
}

func main(){
	lines := readFile("input")
	instructions := lines[0]
	MyGraph := make(map[string]Vertex,0)
	start := make([]Vertex,0)
	for i, line := range lines[2:]{
		vertVal := strings.Trim(strings.Split(line, "=")[0]," ")
		neighs := strings.Split(line,"=")[1]
		neighs = strings.TrimPrefix(strings.TrimSuffix(neighs,")")," (")
		neighVals := strings.Split(neighs,",")
		neighVals[1] = strings.Trim(neighVals[1]," ")
		v := Vertex{id: i,leftNeigh:neighVals[0],rightNeigh:neighVals[1],value:vertVal}
		if v.value[2] == 'A'{
			start = append(start, v)
		}
		MyGraph[v.value] = v
	}
	steps := goTroughGraph1(instructions, MyGraph, "AAA")
	fmt.Printf("Part 1: %d\n", steps)
	fmt.Println(start)
	steps2 := goTroughGraph2(instructions, MyGraph, start)
	fmt.Printf("Part 2: %d\n", steps2)
}