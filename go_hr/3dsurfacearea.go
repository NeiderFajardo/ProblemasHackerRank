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

/*
 * Complete the 'surfaceArea' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY A as parameter.
 */

func surfaceArea(A [][]int32, h, w int32) int32 {
	// Write your code here
	area := int32(0)
	for i := 0; i <= int(h-1); i++ {
		for j := 0; j <= int(w-1); j++ {
			area += (A[i][j] * 4) + 2
			if j >= 1 {
				area -= int32(math.Min(float64(A[i][j-1]), float64(A[i][j]))) * 2
			}
			if i >= 1 {
				area -= int32(math.Min(float64(A[i-1][j]), float64(A[i][j]))) * 2
			}
		}
	}
	return area
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	HTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	H := int32(HTemp)

	WTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	W := int32(WTemp)

	var A [][]int32
	for i := 0; i < int(H); i++ {
		ARowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var ARow []int32
		for _, ARowItem := range ARowTemp {
			AItemTemp, err := strconv.ParseInt(ARowItem, 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			ARow = append(ARow, AItem)
		}

		if len(ARow) != int(W) {
			panic("Bad input")
		}

		A = append(A, ARow)
	}

	result := surfaceArea(A, H, W)

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
