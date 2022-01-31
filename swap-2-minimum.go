package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	var swaps int32
	for i, v := range arr {
		// make index more intuitively match with the proper value
		pos := int32(i) + 1

		// if already placed correctly, no need to swap for any reason
		if v == pos {
			continue
		}

		/*
		 *   Best Case for unordered:
		 *   arr[a] = b AND arr[b] = a
		 *   swap and ordered directly (SOD)
		 */
		// check if SOD could happen
		if arr[v-1] == pos {
			arr[i], arr[v-1] = arr[v-1], arr[i]
			swaps += 1
		} else {
			/*
			 *   Creating best case
			 *   7, 1, 3, ...., 2
			 *   +
			 *   pos = 1, find where 1 is located        (find the ACTUAL number position, get 2)
			 *   find where is 2 located to swap with 7 (find the best case pair with ACTUAL number pos.)
			 *   swap 7 with 2
			 *   swap 2 with 1 (best case reached)
			 */

			// finding the actual number pos (actual number is [pos])
			// finding the actual number pair position (actual number best case pair is actual number pos)
			var actPos int32 = -1
			var actPairPos int32 = -1

			// finding two in one to cut time complexity
			for j := pos; j < int32(len(arr)); j++ {
				if actPos == -1 && arr[j] == pos {
					actPos = j
				} else if actPos != -1 && arr[j] == actPos+1 {
					actPairPos = j
					break
				}
			}

			// if not found yet, find again start with pos till actPos which previously found
			if actPairPos == -1 {
				for j := pos; j < actPos; j++ {
					if arr[j] == actPos+1 {
						actPairPos = j
						break
					}
				}
			}

			// Perform best case
			arr[i], arr[actPairPos] = arr[actPairPos], arr[i]
			arr[i], arr[actPos] = arr[actPos], arr[i]
			swaps += 2
		}
	}
	return swaps
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
