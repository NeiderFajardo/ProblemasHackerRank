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
 * Complete the 'chocolateFeast' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER c
 *  3. INTEGER m
 */

func chocolateFeast(n int32, c int32, m int32) int32 {
    // Write your code here
	waux := int32(0)
	chocolates := n/c
	wrappers := chocolates
	for wrappers >= m {
		waux = wrappers/m
		chocolates += waux
		wrappers = wrappers + waux - (waux*m)
	}
	return chocolates
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)
        c := int32(cTemp)

        mTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
        checkError(err)
        m := int32(mTemp)

        result := chocolateFeast(n, c, m)

        fmt.Printf("%d\n", result)
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
