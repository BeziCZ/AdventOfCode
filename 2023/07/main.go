package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
	"strconv"
	"slices"
	"sort"
)

type Hand struct {
	handId int
	combination string
	handType int
	rank int
	bid int 
	value int
}

var (
	handTypes = map[string]int{
		"FiveOfAKind": 7,
		"FourOfAKind": 6,
		"FullHouse": 5,
		"ThreeOfAKind": 4,
		"TwoPair": 3,
		"Pair": 2,
		"HighCard": 1,
	}
	cardTypes = map[string]string{
		"A": "13",
		"K": "12",
		"Q": "11",
		"J": "00",
		"T": "09",
		"9": "08",
		"8": "07",
		"7": "06",
		"6": "05",
		"5": "04",
		"4": "03",
		"3": "02",
		"2": "01",
	}
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

func getHandType(comb string) int{
	cardOccs := make([]int, 0)
	for cardType, _ := range cardTypes {
		count := strings.Count(comb, cardType)
		if count > 0 {
			cardOccs = append(cardOccs, count)
		}
	}
	switch l := len(cardOccs); {
	case l == 1:
		return 7
	case l == 2 && slices.Max(cardOccs) == 4:
		return 6
	case l == 2 && slices.Max(cardOccs) == 3:
		return 5
	case l == 3 && slices.Max(cardOccs) == 3:
		return 4
	case l == 3 && slices.Max(cardOccs) == 2:
		return 3
	case l == 4:
		return 2
	case l == 5:
		return 1
	}
	return 0
}

func getHandTypeJ(comb string) int{
	cardOccs := make(map[string]int, 0)
	jokerCount := 0
	for cardType, _ := range cardTypes {
		count := strings.Count(comb, cardType)
		if count > 0 {
			if cardType != "J"{
				cardOccs[cardType] = count
			} else {
				jokerCount++
			}
		}
	}
	if comb == "JJJJJ"{
		return getHandType(comb)
	}else if jokerCount != 0{
		maxOcc := 0
		maxKey := ""
		for card, occ := range cardOccs {
			if occ > maxOcc{
				maxKey = card
				maxOcc = occ
			} else if occ == maxOcc{
				m, _ := strconv.Atoi(cardTypes[maxKey])
				c, _ := strconv.Atoi(cardTypes[card])
				if c > m{
					maxKey = card
					maxOcc = occ
				}
			}
		}
		newComb := strings.ReplaceAll(comb, "J", maxKey)
		return getHandType(newComb)
	} else{
		return getHandType(comb)
	}
}

func getHandValue(hand Hand) int{
	valStr := fmt.Sprint(hand.handType)
	for _, card := range hand.combination{
		valStr += cardTypes[string(card)]
	}
	val, _ := strconv.Atoi(valStr)
	return val
}
func main() {
	lines := readFile("./input")

	hands := make([]Hand,0)

	for id, line := range lines{
		split := strings.Split(line, " ")
		b, _ := strconv.Atoi(split[1])
		hand := Hand{combination: split[0], handId: id, handType: getHandTypeJ(split[0]), rank: 0, bid: b}
		hands = append(hands, hand)
	}

	for i, h := range hands{
		hands[i].value = getHandValue(h)
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].value < hands[j].value
	})
	
	for i := 0; i < len(hands); i++{
		hands[i].rank = i+1
	}
	f, _ := os.Create("./sortedinput")
	for _, h := range hands{
		s := "Combination: " + fmt.Sprint(h.combination)
		s += " Hand type: " + fmt.Sprint(h.handType)
		s += " Rank: " + fmt.Sprint(h.rank)
		s += " Bid: " + fmt.Sprint(h.bid) + "\n"
		_, _ = f.WriteString(s)
	}
	total := 0 
	for _, h := range hands{
		total += h.rank * h.bid
	}
	fmt.Println(total)
}