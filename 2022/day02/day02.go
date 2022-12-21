package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	defer file.Close()

	m := map[string]uint64{}
	m["X"] = 1
	m["Y"] = 2
	m["Z"] = 3

	o := map[string]uint64{}
	o["AX"] = 3
	o["AY"] = 6
	o["AZ"] = 0
	o["BX"] = 0
	o["BY"] = 3
	o["BZ"] = 6
	o["CX"] = 6
	o["CY"] = 0
	o["CZ"] = 3

	var accu uint64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			curr := strings.Join(strings.Fields(line), "")
			accu += m[string(curr[1])]
			accu += o[curr]
		}
	}
	fmt.Println(accu)
}

