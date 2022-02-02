package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func moveToNewKey(freq map[int32]map[int32]bool, t int32, b int32, a int32) {
	// b stands for key before and a stands before key after
	// a == 0 if no need to move and only delete

	// delete from previous group
	delete(freq[b], t)
	// add to new group, if willing
	if a != 0 {
		freq[a][t] = true
	}

	// check if current group is empty, delete
	if len(reflect.ValueOf(freq[b]).MapKeys()) == 0 {
		delete(freq, b)
	}
}

func logData(d map[int32]int32, f map[int32]map[int32]bool, before bool) {
	if before {
		fmt.Printf("data before: %v\n", d)
		fmt.Printf("freq before: %v\n", f)
		return
	}
	fmt.Printf("data now: %v\n", d)
	fmt.Printf("freq now: %v\n", f)
}

// Complete the freqQuery function below.
func freqQuery(queries [][]int32) []int32 {
	result := []int32{}
	data := make(map[int32]int32)
	freq := make(map[int32]map[int32]bool)
	for _, q := range queries { // adding data
		fmt.Printf("\ninput now: %v\n", q)
		logData(data, freq, true)
		if q[0] == 1 {
			if _, f := data[q[1]]; f {
				b := data[q[1]]
				data[q[1]] += 1

				// move to new occurence group
				moveToNewKey(freq, q[1], b, b+1)
			} else {
				data[q[1]] = 1
				if _, f := freq[1]; f {
					freq[1][q[1]] = true
				} else {
					freq[1] = map[int32]bool{q[1]: true}
				}
			}
		} else if q[0] == 2 { // remove data
			if _, f := data[q[1]]; f {
				b := data[q[1]]
				data[q[1]] -= 1

				// remove an int if it has 0 frequency
				if data[q[1]] == 0 {
					delete(data, q[1])
				}

				moveToNewKey(freq, q[1], b, b-1)

			}
		} else { // find data whose such given frequency (q[1])
			if _, f := freq[q[1]]; f {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}
		}
		logData(data, freq, false)
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
