package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func gcd(a, b int) int{
	var aux int

	for(a%b != 0){
		aux=a%b
		a = b
		b = aux
	}
	return b
}

// Complete the solve function below.
func solve(a int, b int, c int) string {
	if a > b{
		if c < a && c%(gcd(a,b))== 0{
			return "YES"
		}
	}else{
		if c < b && c%(gcd(b,a))== 0{
			return "YES"
		}
	} 
	return "NO"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        abc := strings.Split(readLine(reader), " ")

        aTemp, err := strconv.ParseInt(abc[0], 10, 64)
        checkError(err)
        a := int(aTemp)

        bTemp, err := strconv.ParseInt(abc[1], 10, 64)
        checkError(err)
        b := int(bTemp)

        cTemp, err := strconv.ParseInt(abc[2], 10, 64)
        checkError(err)
        c := int(cTemp)

        result := solve(a, b, c)

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
