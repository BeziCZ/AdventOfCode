package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

var (
	loop = make(map[int]Pipe, 0)
)
type Pipe struct {
	id int
	pType rune
	prev int
	next int
	Y int
	X int
	
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

func getStart(lines []string) (int, int){
	for i, line := range lines {
		for j, ch := range line{
			if ch == 'S'{
				return i, j
			}
		}
	}
	return 0, 0
}

func checkAroundS(lines []string, PosX int, PosY int, num int) (rune, int, int, string){
	switch num{
	case 0: 
	{
		if lines[PosY-1][PosX] == '7' || lines[PosY-1][PosX] == 'F' || lines[PosY-1][PosX] == '|'{
			return rune(lines[PosY-1][PosX]), PosX, PosY-1, "U"
		}
	}
	case 1:
	{
		if lines[PosY][PosX+1] == '-' || lines[PosY][PosX+1] == 'J' || lines[PosY][PosX+1] == '7'{
			return rune(lines[PosY][PosX+1]), PosX+1, PosY, "R"
		}
	}
	case 2:
	{
		if lines[PosY+1][PosX] == 'J' || lines[PosY+1][PosX] == '|' || lines[PosY+1][PosX] == 'L'{
			return rune(lines[PosY+1][PosX]), PosX, PosY+1, "D"
		}
	}
	case 3:
		if PosX != 0{
			if lines[PosY][PosX-1] == 'L' || lines[PosY][PosX-1] == 'F' || lines[PosY][PosX-1] == '-'{
				return rune(lines[PosY][PosX-1]), PosX-1, PosY, "L"
			}
		}
	}
	return ' ', -1, -1, ""
}

func getNextR(lines []string, X int, Y int, curr Pipe) (string, rune){
	switch curr.pType{
	case '7':
	{
		if rune(lines[Y+1][X]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y+1][X]), Y:Y+1, X:X, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "D", rune(lines[Y+1][X])
	}
	case 'J':
	{
		if rune(lines[Y-1][X]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y-1][X]), Y:Y-1, X:X, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "U", rune(lines[Y-1][X])
	}
	case '-':
	{
		if rune(lines[Y][X+1]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y][X+1]), Y:Y, X:X+1, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "R", rune(lines[Y][X+1])
	}
	}
	return "", ' '
}

func getNextU(lines []string, X int, Y int, curr Pipe) (string, rune){
	switch curr.pType{
	case '|':
	{
		if rune(lines[Y-1][X]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y-1][X]), Y:Y-1, X:X, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "U", rune(lines[Y-1][X])
	}
	case '7':
	{
		if rune(lines[Y][X-1]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y][X-1]), Y:Y, X:X-1, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "L", rune(lines[Y][X-1])
	}
	case 'F':
	{
		if rune(lines[Y][X+1]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y][X+1]), Y:Y, X:X+1, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "R", rune(lines[Y][X+1])
	}
	}
	return "", ' '
}

func getNextD(lines []string, X int, Y int, curr Pipe) (string, rune){
	switch curr.pType{
	case '|':
	{
		if rune(lines[Y+1][X]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y+1][X]), Y:Y+1, X:X, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "D", rune(lines[Y+1][X])
	}
	case 'J':
	{
		if rune(lines[Y][X-1]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y][X-1]), Y:Y, X:X-1, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "L", rune(lines[Y][X-1])
	}
	case 'L':
	{
		if rune(lines[Y][X+1]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y][X+1]), Y:Y, X:X+1, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "R", rune(lines[Y][X+1])
	}
	}
	return "", ' '
}

func getNextL(lines []string, X int, Y int, curr Pipe) (string, rune){
	switch curr.pType{
	case 'F':
	{
		if rune(lines[Y+1][X]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y+1][X]), Y:Y+1, X:X, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "D", rune(lines[Y+1][X])
	}
	case 'L':
	{
		if rune(lines[Y-1][X]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y-1][X]), Y:Y-1, X:X, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "U", rune(lines[Y-1][X])
	}
	case '-':
	{
		if rune(lines[Y][X-1]) ==  'S'{
			if curr, ok := loop[curr.id]; ok {
				curr.next = loop[0].id
				loop[curr.id] = curr
			}
			if start, ok := loop[0]; ok{
				start.prev = curr.id
				loop[0] = start
			}
			return "", 'S'
		}
		next := Pipe{id:curr.id+1, pType: rune(lines[Y][X-1]), Y:Y, X:X-1, prev:curr.id, next:-1}
		if curr, ok := loop[curr.id]; ok {
			curr.next = next.id
			loop[curr.id] = curr
		}
		loop[next.id] = next
		return "L", rune(lines[Y][X-1])
	}
	case 'S':
	{
		if curr, ok := loop[curr.id]; ok {
			curr.next = loop[0].id
			loop[curr.id] = curr
		}
		if start, ok := loop[0]; ok{
			start.prev = curr.id
			loop[0] = start
		}
		return "", 'S'
	}
	}
	return "", ' '
}

func floodFill(node string){

}
func main() {
	lines := readFile("testinput")
	lIdx, idx := getStart(lines)

	start := Pipe{id:0, pType:'S', prev:-1, next:-1, Y:lIdx, X:idx}

	lastMove := ""
	for i := 0; i < 4; i++{
		neighbour, X, Y, move := checkAroundS(lines, start.X, start.Y, i)
		if neighbour != ' '{
			if start.next == -1{
				neig := Pipe{id:1, pType:neighbour, prev: start.id, next: -1, X: X, Y: Y}
				start.next = neig.id
				loop[0] = start
				loop[1] = neig
				lastMove = move
			}
		}
	}
	nextP := ' '
	lastIdx := 1
	for nextP != 'S'{
		switch lastMove{
		case "R":
			lastMove, nextP = getNextR(lines, loop[lastIdx].X, loop[lastIdx].Y, loop[lastIdx])
		case "L":
			lastMove, nextP = getNextL(lines, loop[lastIdx].X, loop[lastIdx].Y, loop[lastIdx])
		case "U":
			lastMove, nextP = getNextU(lines, loop[lastIdx].X, loop[lastIdx].Y, loop[lastIdx])
		case "D":
			lastMove, nextP = getNextD(lines, loop[lastIdx].X, loop[lastIdx].Y, loop[lastIdx])
		}
		lastIdx++
	}

	fmt.Println(len(loop)/2)
	// Visualization. Useless but cute. 
	charMatrix := make([][]string, 0)
	for _, line := range lines {
		repLine := strings.ReplaceAll(line, ".", "O")
		repLine = strings.ReplaceAll(repLine, "F", "┌")
		repLine = strings.ReplaceAll(repLine, "7", "┐")
		repLine = strings.ReplaceAll(repLine, "J", "┘")
		repLine = strings.ReplaceAll(repLine, "L", "└")
		repLine = strings.ReplaceAll(repLine, "-", "─")
		repLine = strings.ReplaceAll(repLine, "|", "│")
		split := strings.Split(repLine, "")
		charMatrix = append(charMatrix, split)
	}

	for _, line := range charMatrix {
		fmt.Println(line)
	}
}
