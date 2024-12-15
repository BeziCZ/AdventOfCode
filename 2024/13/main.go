package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	DiffAX int
	DiffAY int
	DiffBX int
	DiffBY int
	prizeX int
	prizeY int
	currY  int
	currX  int
}

type Result struct {
	Tokens   int
	APresses int
	BPresses int
	Possible bool
}

func readInput(path string) []Position {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var positions []Position
	var pos Position

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			positions = append(positions, pos)
			pos = Position{}
			continue
		}

		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "Button A":
			fmt.Sscanf(parts[1], "X+%d, Y+%d", &pos.DiffAX, &pos.DiffAY)
		case "Button B":
			fmt.Sscanf(parts[1], "X+%d, Y+%d", &pos.DiffBX, &pos.DiffBY)
		case "Prize":
			fmt.Sscanf(parts[1], "X=%d, Y=%d", &pos.prizeX, &pos.prizeY)
		}
	}
	if pos != (Position{}) {
		pos.currX = 0
		pos.currY = 0
		positions = append(positions, pos)
	}
	return positions
}

func gcd(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x1, y1 := gcd(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return gcd, x, y
}

func FindSolutionForDimension(dx1, dx2, target int, maxPresses int) (int, int, bool) {
	// Handle the case where target is 0
	if target == 0 {
		return 0, 0, true
	}

	// Calculate GCD
	gcd, x, y := gcd(dx1, dx2)

	// Check if target is reachable
	if target%gcd != 0 {
		fmt.Printf("Target %d not reachable with moves %d and %d (GCD: %d)\n", target, dx1, dx2, gcd)
		return 0, 0, false
	}

	// Scale the basic solution
	scale := target / gcd
	x *= scale
	y *= scale

	// Adjust solution to minimize 3*x + y while keeping both non-negative
	// We need to handle dx1 and dx2 signs properly
	dx1Div := dx1 / gcd
	dx2Div := dx2 / gcd
	if dx1Div < 0 {
		dx1Div = -dx1Div
	}
	if dx2Div < 0 {
		dx2Div = -dx2Div
	}

	for {
		newX := x - dx2Div
		newY := y + dx1Div

		if newX < 0 || newY >= maxPresses {
			break
		}

		if 3*newX+newY < 3*x+y {
			x = newX
			y = newY
		} else {
			break
		}
	}

	// Make sure solution is positive
	for x < 0 {
		x += dx2Div
		y -= dx1Div
	}

	for y < 0 {
		y += dx1Div
		x -= dx2Div
	}

	// Check limits
	if x >= maxPresses || y >= maxPresses || x < 0 || y < 0 {
		fmt.Printf("Solution exceeds limits: A=%d, B=%d\n", x, y)
		return 0, 0, false
	}

	return x, y, true
}
func FindMinimumTokens(pos Position) Result {
	const maxPresses = 100

	// Convert to relative coordinates
	targetX := pos.prizeX - pos.currX
	targetY := pos.prizeY - pos.currY

	fmt.Printf("Solving for X: target=%d, moves A=%d B=%d\n", targetX, pos.DiffAX, pos.DiffBX)
	aX, bX, possibleX := FindSolutionForDimension(pos.DiffAX, pos.DiffBX, targetX, maxPresses)
	if !possibleX {
		fmt.Println("No solution found for X coordinate")
		return Result{Possible: false}
	}

	fmt.Printf("Solving for Y: target=%d, moves A=%d B=%d\n", targetY, pos.DiffAY, pos.DiffBY)
	aY, bY, possibleY := FindSolutionForDimension(pos.DiffAY, pos.DiffBY, targetY, maxPresses)
	if !possibleY {
		fmt.Println("No solution found for Y coordinate")
		return Result{Possible: false}
	}

	fmt.Printf("Solutions found:\n")
	fmt.Printf("X: %d presses of A, %d presses of B\n", aX, bX)
	fmt.Printf("Y: %d presses of A, %d presses of B\n", aY, bY)

	aPresses := max(aX, aY)
	bPresses := max(bX, bY)

	if aPresses >= maxPresses || bPresses >= maxPresses {
		fmt.Printf("Solution exceeds maximum presses: A=%d, B=%d\n", aPresses, bPresses)
		return Result{Possible: false}
	}

	tokens := 3*aPresses + bPresses

	return Result{
		Tokens:   tokens,
		APresses: aPresses,
		BPresses: bPresses,
		Possible: true,
	}
}

func main() {
	positions := readInput("testinput.txt")
	fmt.Println(positions)
	sum := 0
	for _, pos := range positions {
		result := FindMinimumTokens(pos)
		if result.Possible {
			fmt.Println(result)
			sum += result.Tokens
		}
	}

	fmt.Println(sum)
}
