package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

var program []int
var regA int
var regB int
var regC int
var instrPtr int

func readInput(path string) ([]int, int) {
	file, _ := os.Open(path)
	scanner := bufio.NewScanner(file)
	defer file.Close()

	var regA int
	var program []int

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Register A:") {
			fmt.Sscanf(line, "Register A: %d", &regA)
		} else if strings.HasPrefix(line, "Program:") {
			programStr := strings.TrimPrefix(line, "Program: ")
			programStrs := strings.Split(programStr, ",")
			for _, numStr := range programStrs {
				var num int
				fmt.Sscanf(numStr, "%d", &num)
				program = append(program, num)
			}
		}
	}

	return program, regA
}

func calcCombo(operand int) int {
	switch operand {
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	}
	return 0
}

func adv(operand int) {
	regA = regA / int(math.Pow(float64(2), float64(calcCombo(operand))))
}

func bxl(operand int) {
	regB = regB ^ operand
}

func bst(operand int) {
	regB = calcCombo(operand) % 8
}

func jnz(operand int) bool {
	if regA != 0 {
		instrPtr = operand
		return true
	}
	return false
}

func bxc(operand int) {
	regB = regB ^ regC
}

func out(operand int) string {
	return fmt.Sprintf("%d", calcCombo(operand)%8)
}

func bdv(operand int) {
	regB = regA / int(math.Pow(float64(2), float64(calcCombo(operand))))
}

func cdv(operand int) {
	regC = regA / int(math.Pow(float64(2), float64(calcCombo(operand))))
}

func part1() string {
	outStr := ""
	for {
		if instrPtr >= len(program)-1 {
			break
		}
		instr := program[instrPtr]
		operand := program[instrPtr+1]
		jump := false
		switch instr {
		case 0:
			adv(operand)
		case 1:
			bxl(operand)
		case 2:
			bst(operand)
		case 3:
			jump = jnz(operand)
		case 4:
			bxc(operand)
		case 5:
			if outStr == "" {
				outStr = out(operand)
			} else {
				outStr += ","
				outStr += out(operand)
			}
		case 6:
			bdv(operand)
		case 7:
			cdv(operand)
		}
		if !jump {
			instrPtr += 2
		}
	}
	return outStr
}

func main() {
	program, regA = readInput("testinput.txt")

	// Part 1
	instrPtr = 0
	regB = 0
	regC = 0

	start := time.Now()
	outPart1 := part1()
	elapsed := time.Since(start)
	fmt.Println("Part 1:", outPart1, "Elapsed time:", elapsed)

	programStr := fmt.Sprintf("%d", program[len(program)-1])
	for i := len(program) - 2; i >= 0; i-- {
		programStr += fmt.Sprintf(",%d", program[i])
	}
}
