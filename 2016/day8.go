package main

/*
--- Day 8: Two-Factor Authentication ---

You come across a door implementing what you can only assume is an
implementation of two-factor authentication after a long game of
requirements telephone.

To get past the door, you first swipe a keycard (no problem; there
was one on a nearby desk). Then, it displays a code on a little screen,
and you type that code on a keypad. Then, presumably, the door unlocks.

Unfortunately, the screen has been smashed. After a few minutes, you've
taken everything apart and figured out how it works. Now you just have to
work out what the screen would have displayed.

The magnetic strip on the card you swiped encodes a series of instructions
for the screen; these instructions are your puzzle input. The screen is 50
pixels wide and 6 pixels tall, all of which start off, and is capable of
three somewhat peculiar operations:

rect AxB turns on all of the pixels in a rectangle at the top-left of the
screen which is A wide and B tall.
rotate row y=A by B shifts all of the pixels in row A (0 is the top row)
right by B pixels. Pixels that would fall off the right end appear at the
left end of the row.
rotate column x=A by B shifts all of the pixels in column A
(0 is the left column) down by B pixels. Pixels that would fall off the
bottom appear at the top of the column.
For example, here is a simple sequence on a smaller screen:

rect 3x2 creates a small rectangle in the top-left corner:

###....
###....
.......
rotate column x=1 by 1 rotates the second column down by one pixel:

#.#....
###....
.#.....
rotate row y=0 by 4 rotates the top row right by four pixels:

....#.#
###....
.#.....
rotate column x=1 by 1 again rotates the second column down by one pixel,
causing the bottom pixel to wrap back to the top:

.#..#.#
#.#....
.#.....
As you can see, this display technology is extremely powerful, and will soon
dominate the tiny-code-displaying-screen market. That's what the advertisement
on the back of the display tries to convince you, anyway.

There seems to be an intermediate check of the voltage used by the display:
after you swipe your card, if the screen did work, how many pixels should be lit?

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
	var accu int = 0
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
				accu++
			}
		}
	}
	fmt.Println(accu)
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
