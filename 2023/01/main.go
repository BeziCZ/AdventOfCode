package main

import(
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"io"
	"strings"
)

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func getFirstNumber(s string) string {
	var retVal string = ""
	for _, c := range s{
		if c > '0' && c <= '9' {
			retVal = string(c)
			break
		}
	}
	return retVal
}

func replaceChar(s string, sIndex int, eIndex int, num string) string {
	return(s[0:sIndex] + num + s[eIndex:])
}
func reformatString(s string) string{
	nums := [9]string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
	vals := [9]string{"o1e","t2o","t3e","f4r","f5e","s6x","s7n","e8t","n9e",}
	var tmpStr = ""
	for i := 0; i < len(s); i++{
		tmpStr += string(s[i])
		for idx, num := range nums{
			if strings.Contains(tmpStr, num){
				length := len(num)
				sIndex := strings.Index(tmpStr,num)
				eIndex := sIndex + length
				tmpStr = replaceChar(tmpStr,sIndex,eIndex,vals[idx])
			}
		}
	}
	return tmpStr
}

func main() {
	fmt.Println("Opening file...")
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var sum int = 0
	for {
		line, err := reader.ReadString('\n')
		line = reformatString(line)
		var tmp string = ""
		if err != nil{
			if err == io.EOF {
				tmp = getFirstNumber(line)
				revLine := Reverse(line)
				val := getFirstNumber(revLine)
				tmp += val
				num, _ := strconv.Atoi(tmp)
				sum += num
				break
			}
			log.Fatal(err)
			return
		}
		tmp = getFirstNumber(line)
		revLine := Reverse(line)
		tmp += getFirstNumber(revLine)
		num, err := strconv.Atoi(tmp)
		sum += num
	}
	fmt.Println("The sum is:", sum)
	return
}