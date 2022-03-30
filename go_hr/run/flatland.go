package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the flatlandSpaceStations function below.
func flatlandSpaceStations(n int32, c []int) int {
	var max int = 0
	var actual int = 0
	sort.Ints(c)
	if c[0] != 0 {
		max = c[0]
	}
	if c[len(c)-1] != int(n)-1 {
		if int(n)-c[len(c)-1]-1 > max {
			max = int(n) - c[len(c)-1] - 1
		}
	}
	for i := 0; i < len(c)-1; i++ {
		actual = (c[i+1] - c[i]) / 2
		if max == 0 {
			max = actual
		} else if actual > max {
			max = actual
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int

	for i := 0; i < int(m); i++ {
		cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		cItem := int(cItemTemp)
		c = append(c, cItem)
	}

	result := flatlandSpaceStations(n, c)

	fmt.Printf("%d\n", result)
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
