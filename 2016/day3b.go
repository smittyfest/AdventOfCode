package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 3: Squares With Three Sides ---

--- Part Two ---

Now that you've helpfully marked up their design documents,
it occurs to you that triangles are specified in groups of three
vertically. Each set of three numbers in a column specifies a triangle.
Rows are unrelated.

For example, given the following specification, numbers with the same
hundreds digit would be part of the same triangle:

101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603

In your puzzle input, and instead reading by columns, how many
of the listed triangles are possible?

*/

func main() {
	file, e := os.Open("day3-input.txt")
	if e != nil {
		log.Fatal(e)
	}
	defer file.Close()
	var possibleTriangles int = 0
	// input file is 1734 lines long
	var fields [1734][3]int
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		sides := strings.Fields(scanner.Text())
		sideA, err := strconv.Atoi(sides[0])
		sideB, err := strconv.Atoi(sides[1])
		sideC, err := strconv.Atoi(sides[2])
		if err == nil {
			fields[i][0] = sideA
			fields[i][1] = sideB
			fields[i][2] = sideC
		}
	}
	// 1734 / 3 = 578
	for i := 0; i < 1734; i += 3 {
		for j := 0; j < 3; j++ {
			sideA := fields[i][j]
			sideB := fields[i+1][j]
			sideC := fields[i+2][j]
			if (sideA+sideB) > sideC &&
				(sideB+sideC) > sideA &&
				(sideA+sideC) > sideB {
				possibleTriangles++
			}
		}
	}
	fmt.Printf("The number of possible triangles is %d\n", possibleTriangles)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
