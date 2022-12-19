package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	defer file.Close()

	var max uint64 = 0
	var accu uint64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			i, err := strconv.ParseUint(line, 10, 64)
			if err != nil {
				fmt.Println(err)
        return
			}
			accu += i
		} else {
			if accu > max {
				max = accu
			}
			accu = 0
		}
	}
	fmt.Println(max)
}
