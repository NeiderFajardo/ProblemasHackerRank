package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

/*
 * Complete the 'encryption' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func encryption(s string) string {
	// Write your code here
	l := len(s)
	m := int(math.Sqrt(float64(l)))
	column := m
	result := ""
	if m*m != l {
		column += 1
	}
	for i := 0; i < column; i++ {
		result += string(s[i])
		j := i + column
		for j < len(s) {
			if j < len(s) {
				result += string(s[j])
			}
			j += column
		}
		if i < column-1 {
			result += " "
		}
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	s := readLine(reader)

	result := encryption(s)

	fmt.Println(result)
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
