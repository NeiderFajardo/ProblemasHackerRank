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
 * Complete the 'timeInWords' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER m
 */

func timeInWords(h int32, m int32) string {
	// Write your code here
	hoursName := map[int32]string{
		0:  "",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "quarter",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
		20: "twenty",
		30: "half",
	}

	minWord := "minutes"
	if m == 1 || 60-m == 1 {
		minWord = "minute"
	}
	var time string
	if m == 0 {
		return hoursName[h] + " o' clock"
	}
	if m <= 30 {
		if m == 30 || m == 15 {
			time = hoursName[m] + " past " + hoursName[h]
		}
		if m > 20 && m < 30 {
			fmt.Println(m)
			time = hoursName[20] + " " + hoursName[m-20] + " " + minWord + " past " + hoursName[h]
		}
		if m <= 20 && m != 15 && m != 30 {
			time = hoursName[m] + " " + minWord + " past " + hoursName[h]
		}
		return time
	}
	if m > 30 {
		if 60-m == 15 {
			time = hoursName[60-m] + " to " + hoursName[h+1]
		}
		if 60-m <= 20 && 60-m != 15 {
			fmt.Println(hoursName[60-m])
			time = hoursName[60-m] + " " + minWord + " to " + hoursName[h+1]
		}
		if 60-m > 20 {
			fmt.Println(60 - m)
			time = hoursName[20] + " " + hoursName[(60-m)-20] + " " + minWord + " to " + hoursName[h+1]
		}

	}
	return time
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

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
