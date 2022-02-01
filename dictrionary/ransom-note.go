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
 * Complete the 'checkMagazine' function below.
 *
 * The function accepts following parameters:
 *  1. STRING_ARRAY magazine
 *  2. STRING_ARRAY note
 */

func checkMagazine(magazine []string, note []string) {
	// Write your code here
	magazineMap := make(map[string]int)

	// seed the map
	for _, w := range magazine {
		if _, found := magazineMap[w]; found {
			magazineMap[w] += 1
		} else {
			magazineMap[w] = 1
		}
	}

	// check word availability to creating word from magazine
	for _, w := range note {
		if _, found := magazineMap[w]; found {
			if magazineMap[w] < 1 {
				fmt.Println("No")
				return
			}
			magazineMap[w] -= 1
		} else {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
