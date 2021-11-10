package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'cavityMap' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func cavityMap(grid []string) []string {
	// Write your code here
	var result []string
	var aux string
	result = append(result, grid[0])
	for i:=1; i < len(grid)-1; i++ {
		line := ""
		for pos, char := range grid[i] {
			num := int(char - '0')
			aux= strconv.Itoa(num)
			if pos != 0 && pos < len(grid[i])-1 {
				up, _ := strconv.Atoi(string([]rune(grid[i-1])[pos]))
				left, _ := strconv.Atoi(string([]rune(grid[i])[pos-1]))
				right, _ := strconv.Atoi(string([]rune(grid[i])[pos+1]))
				down, _ := strconv.Atoi(string([]rune(grid[i+1])[pos]))
				if  num > up && num > left && num > right && num > down {
					aux = "X"
				}
			}
			line += aux
		}
		result = append(result, line)
	}
	if len(grid) > 1 {
		result = append(result, grid[len(grid)-1])
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

	for i, resultItem := range result {
		fmt.Printf("%s", resultItem)

		if i != len(result) - 1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
