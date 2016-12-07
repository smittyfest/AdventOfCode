package main

/*
--- Day 6: Signals and Noise ---

Something is jamming your communications with Santa. Fortunately, your signal is only partially jammed, and protocol in situations like this is to switch to a simple repetition code to get the message through.

In this model, the same message is sent repeatedly. You've recorded the repeating message signal (your puzzle input), but the data seems quite corrupted - almost too badly to recover. Almost.

All you need to do is figure out which character is most frequent for each position. For example, suppose you had recorded the following messages:

eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar
The most common character in the first column is e; in the second, a; in the third, s, and so on. Combining these characters returns the error-corrected message, easter.

Given the recording in your puzzle input, what is the error-corrected version of the message being sent?

To begin, get your puzzle input.

Puzzle input:

*/
import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

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
		secret.WriteString(getMostOccurringCharacter(charCounts))
	}
	return secret.String()
}

func getMostOccurringCharacter(charCounts map[string]int) string {
	var max int = 0
	var maxKey = ""
	for key := range charCounts {
		curr := charCounts[key]
		if curr > max {
			max = curr
			maxKey = key
		}
	}
	return maxKey
}
