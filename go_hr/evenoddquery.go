package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func find(x, y int32, arr []int32) string{
	if int(x) < len(arr) && arr[x]== 0 && x!=y{
		return "Odd" 
	}else {
		if arr[x-1]%2==0{
			return "Even"
		}else{
			return "Odd" 
		}

	}
}

// Complete the solve function below.
func solve(arr []int32, queries [][]int32) map[int]string {
	res := make(map[int]string)
	for i, q := range queries{
		res[i] = find(q[0],q[1], arr)
	}
	return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    arrCount, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for arrItr := 0; arrItr < int(arrCount); arrItr++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    var queries [][]int32
    for queriesRowItr := 0; queriesRowItr < int(q); queriesRowItr++ {
        queriesRowTemp := strings.Split(readLine(reader), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != int(2) {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    result := solve(arr, queries)
    for i:=0;i<len(result);i++ {
        fmt.Println(result[i])
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
