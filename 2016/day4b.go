package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
--- Day 4: Security Through Obscurity ---

--- Part Two ---

With all the decoy data out of the way, it's time to decrypt
this list and get moving.

The room names are encrypted by a state-of-the-art shift cipher,
which is nearly unbreakable without the right software. However,
the information kiosk designers at Easter Bunny HQ were not
expecting to deal with a master cryptographer like yourself.

To decrypt a room name, rotate each letter forward through the
alphabet a number of times equal to the room's sector ID. A
becomes B, B becomes C, Z becomes A, and so on. Dashes become spaces.

For example, the real name for qzmt-zixmtkozy-ivhz-343 is very
encrypted name.

What is the sector ID of the room where North Pole objects are stored?

*/

var alphabet = []string{
	"a", "b", "c", "d", "e",
	"f", "g", "h", "i", "j",
	"k", "l", "m", "n", "o",
	"p", "q", "r", "s", "t",
	"u", "v", "w", "x", "y",
	"z",
}
var indices = map[string]int{
	"a": 0, "b": 1, "c": 2, "d": 3, "e": 4,
	"f": 5, "g": 6, "h": 7, "i": 8, "j": 9,
	"k": 10, "l": 11, "m": 12, "n": 13, "o": 14,
	"p": 15, "q": 16, "r": 17, "s": 18, "t": 19,
	"u": 20, "v": 21, "w": 22, "x": 23, "y": 24,
	"z": 25,
}

func main() {
	file, err := os.Open("day4-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var northPoleSectorId int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		length := len(line)
		checksum := line[length-6 : length-1]
		sector := line[length-10 : length-7]
		encryptedName := strings.Replace(line[0:length-11], "-", "", -1)
		var charCounts = map[string]int{}
		for i := 0; i < len(encryptedName); i++ {
			name := string(encryptedName[i])
			if charCounts[name] > 0 {
				charCounts[name] = charCounts[name] + 1
			} else {
				charCounts[name] = 1
			}
		}
		orderedKeys := getOrderingByOccurrence(charCounts)
		if validateRoom(charCounts, orderedKeys, checksum) {
			sectorVal, err := strconv.Atoi(sector)
			if err == nil {
				decrypted := decrypt(encryptedName, sectorVal)
				if strings.Contains(decrypted, "north") {
					northPoleSectorId = sectorVal
				}
			}
		}
	}
	fmt.Printf("The sector ID of the room where North Pole objects are stored is %d\n", northPoleSectorId)
}

/*
 * returns a reverse-sorted (maximum-first) array of integers
 * representing the number of occurrences of each character
 * in a string.
 */
func getOrderingByOccurrence(charCounts map[string]int) []int {
	var ordered []int
	for key := range charCounts {
		ordered = append(ordered, charCounts[key])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ordered)))
	return ordered
}

/*
 * validates whether a room is "real" or not by determining if the
 * checksum is in the correct format and that the encrypted value
 * of the room is compliant with the checksum.
 */
func validateRoom(charCounts map[string]int, ordered []int, checksum string) bool {
	for i := 0; i < len(checksum); i++ {
		current := string(checksum[i])
		if charCounts[current] != ordered[i] {
			return false
		}
		if i < len(checksum)-1 {
			next := string(checksum[i+1])
			if charCounts[current] == charCounts[next] {
				if current > next {
					return false
				}
			}
		}
	}
	return true
}

func decrypt(encryptedName string, sectorId int) string {
	var decryptedName bytes.Buffer
	for i := 0; i < len(encryptedName); i++ {
		current := string(encryptedName[i])
		if current == "-" {
			decryptedName.WriteString(" ")
		} else {
			newIndex := (indices[current] + sectorId) % 26
			decryptedName.WriteString(alphabet[newIndex])
		}
	}
	return decryptedName.String()
}
