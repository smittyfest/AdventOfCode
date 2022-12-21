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

	var accu uint64 = 0
	var maxCalories1 uint64 = 0
	var maxCalories2 uint64 = 0
	var maxCalories3 uint64 = 0

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
			if accu > maxCalories3 {
				maxCalories3 = accu
			}
			if maxCalories3 > maxCalories2 {
				maxCalories3, maxCalories2 = maxCalories2, maxCalories3
			}
			if maxCalories2 > maxCalories1 {
				maxCalories2, maxCalories1 = maxCalories1, maxCalories2
			}
			accu = 0
		}
	}
	fmt.Println(maxCalories1 + maxCalories2 + maxCalories3)
}
