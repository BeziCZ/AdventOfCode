package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"log"
	"slices"
	"strconv"
)

type Lens struct {
	label string
	strength int
}
type Box struct{
	lenses []Lens // lenses present numOfLens : location
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	raw, err := ioutil.ReadAll(file)
	line := strings.Trim(string(raw),"\n")
	strs := strings.Split(line,",")
	//hashSum := 0
	totalFocusVal := 0
	boxes := make(map[int]Box,0)
	for _, s := range strs{
		currHash := 0
		operIdx := -1
		var currString []string
		if strings.Contains(s, "=") {
			currString = strings.Split(s,"=")
			operIdx = strings.Index(s,"=")
		} else{
			currString = strings.Split(s,"-")
			operIdx = strings.Index(s,"-")
		}
		for i := 0; i < 2; i++ {
			currHash = (17*(currHash + int(s[i])))%256
		}
		if _, ok := boxes[currHash]; ok{
			box := boxes[currHash]
			if string(s[operIdx]) == "-"{
				for _, lens := range box.lenses{
					if lens.label == (currString[0]){
						idx := slices.Index(box.lenses, lens)
						if idx != -1 {
							box.lenses = append(box.lenses[:idx], box.lenses[idx+1:]...)
						}
					}
				}
			} else {
				wasInside := false
				for _, lens := range box.lenses{
					if lens.label == currString[0]{
						wasInside = true
						newLens := lens
						newLens.strength, _ = strconv.Atoi(currString[1])
						idx := slices.Index(box.lenses, lens)
						box.lenses[idx] = newLens
					}
				}
				if !wasInside{
					stren, _ := strconv.Atoi(currString[1])
					newLens := Lens{label: currString[0], strength:stren}
					box.lenses = append(box.lenses, newLens)
				}
			}
			boxes[currHash] = box
		} else{
			newLabel := currString[0]
			newStregth, _ := strconv.Atoi(currString[1])
			newLens := Lens{label: newLabel, strength: newStregth}
			newLenses := []Lens{newLens}
			box := Box{lenses: newLenses}
			boxes[currHash] = box
		}
		/* Part 1
		for _,c := range s{
			currHash = (17*(currHash + int(c)))%256
		}
		*/
		//hashSum += currHash // Pt1
	}
	fmt.Println(boxes)
	for num, box := range boxes{
		currFocus := 0
		for i, lens := range box.lenses{
			currFocus += (1+num)*(i+1)*lens.strength
		}
		totalFocusVal += currFocus
	}
	fmt.Println(totalFocusVal)
	//fmt.Println(hashSum)
}