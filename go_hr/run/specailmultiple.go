package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func toBin(num int)int{
	n := 1
	res := 0
	for n*2 <= num{
		n = n*2
	}
	for n >= 1{
		if res == 0{
			if num/n != 0{
				res = res + num/n
			}
		}else{
			if num/n == 0{
				res = res*10
			}else{
				res = (res*10)+1
			}
		}
		num = num % n
		n = n / 2 

	}
	return res
}

// Complete the solve function below.
func solve(n int) string {
	res := 1
	for (toBin(res)*9)%n != 0{
		res++
	}

	return strconv.Itoa(toBin(res)*9)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int(nTemp)

        result := solve(n)

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
