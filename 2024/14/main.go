package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

type Robot struct {
	posX int
	posY int
	velX int
	velY int
}

//var mapSizeX int = 11
//var mapSizeY int = 7

var mapSizeX = 101
var mapSizeY = 103

func readInput(path string) []Robot {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var robots []Robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		robot := Robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &robot.posX, &robot.posY, &robot.velX, &robot.velY)
		robots = append(robots, robot)
	}
	return robots
}

func moveHorizontal(robot Robot) int {
	newPosX := robot.posX + robot.velX
	if newPosX < 0 {
		newPosX = ((newPosX % mapSizeX) + mapSizeX) % mapSizeX
	} else {
		newPosX = newPosX % mapSizeX
	}
	return newPosX
}

func moveVertical(robot Robot) int {
	newPosY := robot.posY + robot.velY
	// Wrap around using modulo
	if newPosY < 0 {
		newPosY = ((newPosY % mapSizeY) + mapSizeY) % mapSizeY
	} else {
		newPosY = newPosY % mapSizeY
	}
	return newPosY
}

func printMap(robots []Robot) {
	bathroomMap := make([][]string, mapSizeY)
	for i := range bathroomMap {
		bathroomMap[i] = make([]string, mapSizeX)
		for j := range bathroomMap[i] {
			bathroomMap[i][j] = "."
		}
	}
	for _, robot := range robots {
		bathroomMap[robot.posY][robot.posX] = "A"
	}
	for _, row := range bathroomMap {
		fmt.Println(row)
	}
}
func moveRobots(robots []Robot, seconds int) []Robot {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for i := 0; i < seconds; i++ {
		for j := 0; j < len(robots); j++ {
			robots[j].posX = moveHorizontal(robots[j])
			robots[j].posY = moveVertical(robots[j])
		}
		// Extension for part 2
		fmt.Println("Second:", i+1)
		printMap(robots)
		for {
			_, key, err := keyboard.GetKey()
			if err != nil {
				panic(err)
			}
			if key == keyboard.KeySpace {
				break
			}
		}
	}
	return robots
}

func main() {
	robots := readInput("input.txt")
	start := time.Now()
	moveRobots(robots, 10000000)
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, robot := range robots {
		if robot.posX < mapSizeX/2 && robot.posY < mapSizeY/2 {
			q1++
		}
		if robot.posX > mapSizeX/2 && robot.posY < mapSizeY/2 {
			q2++
		}
		if robot.posX < mapSizeX/2 && robot.posY > mapSizeY/2 {
			q3++
		}
		if robot.posX > mapSizeX/2 && robot.posY > mapSizeY/2 {
			q4++
		}
	}
	total := q1 * q2 * q3 * q4
	elapsed := time.Since(start)
	fmt.Println("Part 1:", total, "Time:", elapsed)
}
