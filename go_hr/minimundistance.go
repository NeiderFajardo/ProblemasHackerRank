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
 * Complete the 'minimumDistances' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func minimumDistances(a []int32) int32 {
	// Write your code here
	numbers := make(map[int32]int32)
	min := int32(0)
	for i := 0; i < len(a); i++ {
		if _, ok := numbers[a[i]]; ok {
			if min == 0 {
				min = int32(i) - numbers[a[i]]
			} else {
				if int32(i)-numbers[a[i]] < min {
					min = int32(i) - numbers[a[i]]
				}
			}
			numbers[a[i]] = int32(i)
		} else {
			numbers[a[i]] = int32(i)
		}
	}
	if min == 0 {
		return int32(-1)
	}
	return min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := minimumDistances(a)

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
