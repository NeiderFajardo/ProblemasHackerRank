package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func masCorto(arr []int) int {
	min := arr[0]
	for _, a := range arr {
		if a < min {
			min = a
		}
	}
	return min
}

func eliminarValor(arr []int) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			result = append(result, arr[i])
		}
	}

	return result
}

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int) []int {
	var res []int
	min := masCorto(arr)

	for len(arr) > 0 {
		minaux := 100000
		if len(arr) > 1 {
			res = append(res, len(arr))
			for i, x := range arr {
				if x > min {
					arr[i] = x - min
					if x-min < minaux {
						minaux = x - min
					}
				} else {
					arr[i] = 0
				}
			}
			arr = eliminarValor(arr)
			min = minaux
		} else {
			res = append(res, 1)
			arr = nil
		}
		fmt.Println(arr)
	}
	return res
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := cutTheSticks(arr)

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
