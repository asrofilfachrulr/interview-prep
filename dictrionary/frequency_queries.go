package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	result := []int32{}
	data := make(map[int32]int32)
	for _, q := range queries { // adding data
		if q[0] == 1 {
			if _, f := data[q[1]]; f {
				data[q[1]] += 1
			} else {
				data[q[1]] = 1
			}
		} else if q[0] == 2 { // remove data
			if _, f := data[q[1]]; f {
				data[q[1]] -= 1

				// remove an int if it has 0 frequency
				if data[q[1]] == 0 {
					delete(data, q[1])
				}
			}
		} else { // find data whose such given frequency (q[1])
			isFounded := false
			for d := range data {
				if data[d] == q[1] {
					result = append(result, 1)
					isFounded = true
					break
				}
			}
			if !isFounded {
				result = append(result, 0)
			}
		}
	}
	return result
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

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
