package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'taumBday' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER b
 *  2. INTEGER w
 *  3. INTEGER bc
 *  4. INTEGER wc
 *  5. INTEGER z
 */

func taumBday(b int64, w int64, bc int64, wc int64, z int64) int64 {
	// Write your code here
	var bx int64 = -1
	var wx int64 = -1
	var result *big.Int
	if (bc + z) <= wc {
		bx = bc + z
	} else if (wc + z) <= bc {
		wx = wc + z
	}

	if bx >= 0 {
		result = big.NewInt((b * bc) + (w * bx))
	} else if wx >= 0 {
		result = big.NewInt((b * wx) + (w * wc))
	} else {
		result = big.NewInt((b * bc) + (w * wc))
	}
	return result.Int64()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		bTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		b := int64(bTemp)

		wTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		w := int64(wTemp)

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		bcTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		bc := int64(bcTemp)

		wcTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
		checkError(err)
		wc := int64(wcTemp)

		zTemp, err := strconv.ParseInt(secondMultipleInput[2], 10, 64)
		checkError(err)
		z := int64(zTemp)

		result := taumBday(b, w, bc, wc, z)

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
