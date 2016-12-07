package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

/*
--- Day 6: Signals and Noise ---

--- Part Two ---

Of course, that would be the message - if you hadn't agreed to use
a modified repetition code instead.

In this modified code, the sender instead transmits what looks like
random data, but for each character, the character they actually want
to send is slightly less likely than the others. Even after signal-jamming
noise, you can look at the letter distributions in each column and choose
the least common letter to reconstruct the original message.

In the above example, the least common character in the first column is a;
in the second, d, and so on. Repeating this process for the remaining
characters produces the original message, advent.

Given the recording in your puzzle input and this new decoding methodology,
what is the original message that Santa is trying to send?

*/

func main() {
	file, err := os.Open("day6-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	const rows = 546
	const columns = 8
	var chars [rows][columns]string
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j := 0; j < columns; j++ {
			chars[i][j] = string(line[j])
		}
	}
	secretMessage := decodeMessage(chars, rows, columns)
	fmt.Printf("The secret message is %s\n", secretMessage)
}

func decodeMessage(chars [546][8]string, rows int, columns int) string {
	var secret bytes.Buffer
	for j := 0; j < columns; j++ {
		var charCounts = map[string]int{}
		for i := 0; i < rows; i++ {
			// Depth-First Traverse down each column
			ch := string(chars[i][j])
			if charCounts[ch] > 0 {
				charCounts[ch] = charCounts[ch] + 1
			} else {
				charCounts[ch] = 1
			}
		}
		secret.WriteString(getLeastOccurringCharacter(charCounts))
	}
	return secret.String()
}

func getLeastOccurringCharacter(charCounts map[string]int) string {
	var min int = 1000000
	var minKey = ""
	for key := range charCounts {
		curr := charCounts[key]
		if curr < min {
			min = curr
			minKey = key
		}
	}
	return minKey
}
