package main

import (
	"fmt"
	"strconv"
	"time"
)

//var testinput []int = []int{125, 17}

var input []int = []int{64599, 31, 674832, 2659361, 1, 0, 8867, 321}

// Part 1

type Node struct {
	stone int
	left  *Node
	right *Node
}

func buildTree(stone int, depth int) *Node {
	if depth >= 26 {
		return nil
	}

	node := &Node{stone: stone}

	if stone == 0 {
		if depth < 24 {
			node.left = buildTree(1, depth+1)
		}
		return node
	}

	strval := strconv.Itoa(stone)
	if len(strval)%2 == 0 {
		mid := len(strval) / 2

		left_val, _ := strconv.Atoi(strval[:mid])
		right_val, _ := strconv.Atoi(strval[mid:])

		node.left = buildTree(left_val, depth+1)
		node.right = buildTree(right_val, depth+1)
	} else {
		node.left = buildTree(node.stone*2024, depth+1)
	}

	return node
}

func countLeafNodes(root *Node) int {
	if root == nil {
		return 0
	}

	// If node is a leaf (no children), return 1
	if root.left == nil && root.right == nil {
		return 1
	}

	// Return sum of leaf nodes in left and right subtrees
	return countLeafNodes(root.left) + countLeafNodes(root.right)
}

// Part 2
type Tuple struct {
	stone int
	blink int
}

var mem = make(map[Tuple]int)

func solvePart2(stone int, blinks int) int {
	var val int
	if blinks == 0 {
		return 1
	} else if _, exists := mem[Tuple{stone, blinks}]; exists {
		return mem[Tuple{stone, blinks}]
	} else if stone == 0 {
		val = solvePart2(1, blinks-1)
	} else if len(strconv.Itoa(stone))%2 == 0 {
		strStone := strconv.Itoa(stone)
		mid := len(strconv.Itoa(stone)) / 2
		leftStone, _ := strconv.Atoi(strStone[:mid])
		rightStone, _ := strconv.Atoi(strStone[mid:])
		val = solvePart2(leftStone, blinks-1) + solvePart2(rightStone, blinks-1)
	} else {
		val = solvePart2(stone*2024, blinks-1)
	}
	mem[Tuple{stone, blinks}] = val
	return val
}

func main() {

	total := 0
	start := time.Now()
	for _, stone := range input {
		root := buildTree(stone, 0)
		total += countLeafNodes(root)
	}
	elapsed := time.Since(start)
	fmt.Println("Part 1:", total, "in time:", elapsed)

	total2 := 0
	start = time.Now()
	for _, stone := range input {
		total2 += solvePart2(stone, 75)
	}
	elapsed = time.Since(start)
	fmt.Println("Part 2:", total2, "in time:", elapsed)
}
