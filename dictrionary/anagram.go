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
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	// fmt.Println("\nstring now", s)
	var num int32 = 0
	// map all of substring and group them based on their length
	lenSubstringsMap := make(map[int32][]string)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i; j++ {
			if i == 0 && j == 0 {
				continue
			}
			w := s[i : len(s)-j]
			l := int32(len(w))
			lenSubstringsMap[l] = append(lenSubstringsMap[l], w)
		}
	}
	// fmt.Println(lenSubstringsMap)

	// check anagrams every substrings that has same length
	for l := range lenSubstringsMap {
		// check number of anagram
		num += howMuchAnagram(lenSubstringsMap[l])
	}

	return num
}

func howMuchAnagram(substrings []string) int32 {
	var num int32 = 0
	for i := 0; i < len(substrings)-1; i++ {
		s := substrings[i]
		// fmt.Println("s now: ", s)
		if i == len(substrings)-1 {
			break
		}
		// create temp map for chars of s
		charMap := make(map[string]int32)
		for _, c := range s {
			ch := string(c)
			if _, f := charMap[ch]; f {
				charMap[ch] += 1
			} else {
				charMap[ch] = 1
			}
		}
		// fmt.Println(charMap)
		// check if every word meet, has identical composite / anagram
		for j := i + 1; j < len(substrings); j++ {
			tempMap := make(map[string]int32)
			for k := range charMap {
				tempMap[k] = charMap[k]
			}

			// fmt.Println("comparing with :", substrings[j])
			// fmt.Println(tempMap)
			for _, c := range substrings[j] {
				ch := string(c)
				if _, f := tempMap[ch]; f {
					if tempMap[ch] == 0 {
						// has not enough pair
						break
					}
					tempMap[ch] -= 1
				} else {
					// char has not pair
					break
				}
			}
			// check for every key (a.k.a char) in charMap must be zero means has its pair
			isZero := true
			for k := range charMap {
				if tempMap[k] != 0 {
					isZero = false
					break
				}
			}
			if isZero {
				num += 1
			}
			// fmt.Println(tempMap, "\n")
		}
	}
	return num
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
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
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
