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
 * Complete the 'twoStrings' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func twoStrings(s1 string, s2 string) string {
	// Find the shorter and longer string
	var l string
	var s string
	if len(s1) <= len(s2) {
		s = s1
		l = s2
	} else {
		s = s2
		l = s1
	}

	/*
	 *   spread the longest string to its chars into a dictionary
	 *   then, iterate shorter string chars and map into the dictionary
	 *   return true if the char exists in dictionary
	 */

	dictOfL := make(map[rune]bool)
	for _, c := range l {
		dictOfL[c] = true
	}

	var status string = "NO"

	for _, c := range s {
		if _, exist := dictOfL[c]; exist {
			return "YES"
		}
	}

	return status
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s1 := readLine(reader)

		s2 := readLine(reader)

		result := twoStrings(s1, s2)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
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
