package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

// 1 2 3 4
// 3 2 1 4 => bribes = 2, bribeTwice = 1
// [1] moved 2 backward => bribed by 2 people
// [3] moved 2 forward => bribed twice

func minimumBribes(q []int32) {
	// Write your code here
	var bribes int32 = 0

	for i := len(q) - 1; i >= 0; i-- {
		pos := int32(i) + 1
		v := q[i]

		// check if v placed 2 forward
		if pos < v-2 {
			fmt.Println("Too chaotic")
			return
		}

		// check if number bribed indicated by located backward
		j := int(math.Max(float64(v)-2, 0))
		for ; j < i; j++ {
			if q[j] > v {
				bribes += 1
			}
		}

	}

	fmt.Println(bribes)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
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
