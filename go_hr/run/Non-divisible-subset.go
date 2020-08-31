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
 * Complete the 'nonDivisibleSubset' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY s
 */

func mod2() func(int, int) int {
	res := func(a, b int) int {
		if a >= b {
			return a % b
		}
		return a
	}
	return res
}

func nonDivisibleSubset(k int, s []int) int {
	sumas := make(map[int]int)
	result := 0
	modulo := mod2()
	if k == 1 || k == 0 {
		return 1
	}
	for _, num := range s {
		if _, ok := sumas[modulo(num, k)]; ok {
			sumas[modulo(num, k)] = sumas[modulo(num, k)] + 1
		} else {
			sumas[modulo(num, k)] = 1
		}
	}
	for i, val := range sumas {
		if val != 0 {
			if i != 0 {
				if sum, ok := sumas[k-i]; ok {
					if i != k-i {
						if sum > val {
							result = result + sum
						} else {
							result = result + val
						}
					} else {
						result = result + 1
					}
					sumas[k-i] = 0
				} else {
					result = result + val
				}
			} else {
				result = result + 1
			}
		}
	}
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int(kTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int(sItemTemp)
		s = append(s, sItem)
	}

	result := nonDivisibleSubset(k, s)

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
