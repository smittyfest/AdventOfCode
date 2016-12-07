package main

/*
--- Day 7: Internet Protocol Version 7 ---

While snooping around the local network of EBHQ, you compile a list of IP addresses (they're IPv7, of course; IPv6 is much too limited). You'd like to figure out which IPs support TLS (transport-layer snooping).

An IP supports TLS if it has an Autonomous Bridge Bypass Annotation, or ABBA. An ABBA is any four-character sequence which consists of a pair of two different characters followed by the reverse of that pair, such as xyyx or abba. However, the IP also must not have an ABBA within any hypernet sequences, which are contained by square brackets.

For example:

abba[mnop]qrst supports TLS (abba outside square brackets).
abcd[bddb]xyyx does not support TLS (bddb is within square brackets, even though xyyx is outside square brackets).
aaaa[qwer]tyui does not support TLS (aaaa is invalid; the interior characters must be different).
ioxxoj[asdfgh]zxcvbn supports TLS (oxxo is outside square brackets, even though it's within a larger string).
How many IPs in your puzzle input support TLS?

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
