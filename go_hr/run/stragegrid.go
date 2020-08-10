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
 * Complete the strangeGrid function below.
 */
func strangeGrid(r int, c int) int {
    /*
     * Write your code here.
     */
	var res int
	par := 1
	impar := 0
	for i := 1; i < c; i++{
		par += 2
		impar += 2
	}
	if r%2 == 0{
		aux := ((r/2)-1)*10
		res = par+aux
	}else{
		aux := ((r/2))*10
		res = impar+aux
	}


	return res
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    rc := strings.Split(readLine(reader), " ")

    rTemp, err := strconv.ParseInt(rc[0], 10, 64)
    checkError(err)
    r := int(rTemp)

    cTemp, err := strconv.ParseInt(rc[1], 10, 64)
    checkError(err)
    c := int(cTemp)

    result := strangeGrid(r, c)

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
