package main

/*
--- Day 8: Two-Factor Authentication ---

--- Part Two ---

You notice that the screen is only capable of displaying capital letters; in the font it uses, each letter is 5 pixels wide and 6 tall.

After you swipe your card, what code is the screen trying to display?

*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day8-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	screen := [6][50]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "rect") {
			instruction := strings.Split(line, " ")
			screen = applyRectangle(screen, instruction)
		} else {
			instruction := strings.Split(line, " ")
			if instruction[1] == "row" {
				screen = rotateRow(screen, instruction)
			} else {
				screen = rotateColumn(screen, instruction)
			}
		}
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < 50; j++ {
			if screen[i][j] == 1 {
				fmt.Printf("X")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func applyRectangle(screen [6][50]int, instruction []string) [6][50]int {
	dimensions := strings.Split(instruction[1], "x")
	a := dimensions[0]
	b := dimensions[1]
	columns, e1 := strconv.Atoi(a)
	rows, e2 := strconv.Atoi(b)
	if e1 == nil && e2 == nil {
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				screen[i][j] = 1
			}
		}
	}
	return screen
}

func rotateRow(screen [6][50]int, instruction []string) [6][50]int {
	dimensions := strings.Split(instruction[2], "=")
	row, e1 := strconv.Atoi(dimensions[1])
	shift, e2 := strconv.Atoi(instruction[4])
	if e1 == nil && e2 == nil {
		oldValues := [50]int{}
		for j := 0; j < 50; j++ {
			oldValues[j] = screen[row][j]
		}
		for j := 0; j < 50; j++ {
			newIndex := (j + shift) % 50
			screen[row][newIndex] = oldValues[j]
		}
	}
	return screen
}

func rotateColumn(screen [6][50]int, instruction []string) [6][50]int {
	dimensions := strings.Split(instruction[2], "=")
	column, e1 := strconv.Atoi(dimensions[1])
	shift, e2 := strconv.Atoi(instruction[4])
	if e1 == nil && e2 == nil {
		oldValues := [6]int{}
		for i := 0; i < 6; i++ {
			oldValues[i] = screen[i][column]
		}
		for i := 0; i < 6; i++ {
			newIndex := (i + shift) % 6
			screen[newIndex][column] = oldValues[i]
		}
	}
	return screen
}
