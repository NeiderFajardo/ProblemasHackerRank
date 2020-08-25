package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func gcd(a, b int) int {
	var aux int

	if b == 0 {
		return a
	}
	for a%b != 0 {
		aux = a % b
		a = b
		b = aux
	}
	return b
}

// Complete the solve function below.
func solve(x1, y1, x2, y2 int) float64 {
	n := gcd(y2-y1, x2-x1)
	return math.Abs(float64(n)) - 1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		x1Y1X2Y2 := strings.Split(readLine(reader), " ")

		x1Temp, err := strconv.ParseInt(x1Y1X2Y2[0], 10, 64)
		checkError(err)
		x1 := int(x1Temp)

		y1Temp, err := strconv.ParseInt(x1Y1X2Y2[1], 10, 64)
		checkError(err)
		y1 := int(y1Temp)

		x2Temp, err := strconv.ParseInt(x1Y1X2Y2[2], 10, 64)
		checkError(err)
		x2 := int(x2Temp)

		y2Temp, err := strconv.ParseInt(x1Y1X2Y2[3], 10, 64)
		checkError(err)
		y2 := int(y2Temp)

		result := solve(x1, y1, x2, y2)

		fmt.Println(result)
	}

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
