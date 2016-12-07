package main

/*
--- Day 7: Internet Protocol Version 7 ---

--- Part Two ---

You would also like to know which IPs support SSL (super-secret listening).

An IP supports SSL if it has an Area-Broadcast Accessor, or ABA, anywhere in the supernet sequences (outside any square bracketed sections), and a corresponding Byte Allocation Block, or BAB, anywhere in the hypernet sequences. An ABA is any three-character sequence which consists of the same character twice with a different character between them, such as xyx or aba. A corresponding BAB is the same characters but in reversed positions: yxy and bab, respectively.

For example:

aba[bab]xyz supports SSL (aba outside square brackets with corresponding bab within square brackets).
xyx[xyx]xyx does not support SSL (xyx, but no corresponding yxy).
aaa[kek]eke supports SSL (eke in supernet with corresponding kek in hypernet; the aaa sequence is not related, because the interior character must be different).
zazbz[bzb]cdb supports SSL (zaz has no corresponding aza, but zbz has a corresponding bzb, even though zaz and zbz overlap).
How many IPs in your puzzle input support SSL?

*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("day7-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var accu int = 0
	for scanner.Scan() {
		line := scanner.Text()
		var withinBrackets bool = false
		var found bool = false
		for i := 0; i <= len(line)-4; i++ {
			if string(line[i]) == "[" {
				withinBrackets = true

			}
			if string(line[i]) == "]" {
				withinBrackets = false
			}
			if string(line[i]) == string(line[i+3]) &&
				string(line[i+1]) == string(line[i+2]) {
				if withinBrackets {
					if string(line[i]) != string(line[i+1]) {
						found = false
						break
					}
				}
				if string(line[i]) == string(line[i+1]) {
					found = false
				} else {
					found = true
				}
			}
		}
		if found {
			accu++
		}
	}
	fmt.Printf("The number of IP lines supporting tls is %d\n", accu)
}
