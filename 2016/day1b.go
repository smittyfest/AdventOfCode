package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
--- Day 1: No Time for a Taxicab ---

--- Part Two ---

Then, you notice the instructions continue on the back of the Recruiting Document. Easter Bunny HQ is actually at the first location you visit twice.

For example, if your instructions are R8, R4, R4, R8, the first location you visit twice is 4 blocks away, due East.

How many blocks away is the first location you visit twice?

Puzzle input:
R4, R1, L2, R1, L1, L1, R1, L5, R1, R5, L2, R3, L3, L4, R4, R4,
R3, L5, L1, R5, R3, L4, R1, R5, L1, R3, L2, R3, R1, L4, L1, R1,
L1, L5, R1, L2, R2, L3, L5, R1, R5, L1, R188, L3, R2, R52, R5,
L3, R79, L1, R5, R186, R2, R1, L3, L5, L2, R2, R4, R5, R5, L5,
L4, R5, R3, L4, R4, L4, L4, R5, L4, L3, L1, L4, R1, R2, L5, R3,
L4, R3, L3, L5, R1, R1, L3, R2, R1, R2, R2, L4, R5, R1, R3, R2,
L2, L2, L1, R2, L1, L3, R5, R1, R4, R5, R2, R2, R4, R4, R1, L3,
R4, L2, R2, R1, R3, L5, R5, R2, R5, L1, R2, R4, L1, R5, L3, L3,
R1, L4, R2, L2, R1, L1, R4, R3, L2, L3, R3, L2, R1, L4, R5, L1,
R5, L2, L1, L5, L2, L5, L2, L4, L2, R3
*/

const (
	NORTH = "NORTH"
	SOUTH = "SOUTH"
	EAST  = "EAST"
	WEST  = "WEST"
)

var current_orientation = NORTH
var visited = make(map[string]bool)

func main() {
	input := "R4, R1, L2, R1, L1, L1, R1, L5, R1, R5, L2, R3, L3, L4, R4, R4, R3, L5, L1, R5, R3, L4, R1, R5, L1, R3, L2, R3, R1, L4, L1, R1, L1, L5, R1, L2, R2, L3, L5, R1, R5, L1, R188, L3, R2, R52, R5, L3, R79, L1, R5, R186, R2, R1, L3, L5, L2, R2, R4, R5, R5, L5, L4, R5, R3, L4, R4, L4, L4, R5, L4, L3, L1, L4, R1, R2, L5, R3, L4, R3, L3, L5, R1, R1, L3, R2, R1, R2, R2, L4, R5, R1, R3, R2, L2, L2, L1, R2, L1, L3, R5, R1, R4, R5, R2, R2, R4, R4, R1, L3, R4, L2, R2, R1, R3, L5, R5, R2, R5, L1, R2, R4, L1, R5, L3, L3, R1, L4, R2, L2, R1, L1, R4, R3, L2, L3, R3, L2, R1, L4, R5, L1, R5, L2, L1, L5, L2, L5, L2, L4, L2, R3"
	distance := distanceToBunnyHeadquarters(input)
	fmt.Printf("The Distance to Bunny HQ is %d\n", distance)
}

func distanceToBunnyHeadquarters(input string) int {
	var curr_x, curr_y int = 0, 0
	var found bool = false
	visited["00"] = true
	coordinates := strings.Split(input, ", ")
	for i := 0; i < len(coordinates); i++ {
		dir, spaces := getSpecs(coordinates[i])
		if dir == "L" {
			found, curr_x, curr_y = getCoordinatesForMoveLeft(spaces, curr_x, curr_y)
		} else {
			found, curr_x, curr_y = getCoordinatesForMoveRight(spaces, curr_x, curr_y)
		}
		if found {
			return int(math.Abs(float64(curr_x + curr_y)))
		}
	}
	return int(math.Abs(float64(curr_x + curr_y)))
}

func getSpecs(directions string) (string, int) {
	dir := string(directions[0:1])
	spaces := string(directions[1:])
	moves, err := strconv.Atoi(spaces)
	if err == nil {
		return dir, moves
	} else {
		panic(err)
	}
}

func getCoordinatesForMoveLeft(left_spaces int, x int, y int) (bool, int, int) {
	switch current_orientation {
	case NORTH:
		current_orientation = WEST
		return updateX(visited, x, x-left_spaces, y)
	case SOUTH:
		current_orientation = EAST
		return updateX(visited, x, x+left_spaces, y)
	case EAST:
		current_orientation = NORTH
		return updateY(visited, x, y, y+left_spaces)
	case WEST:
		current_orientation = SOUTH
		return updateY(visited, x, y, y-left_spaces)
	default:
		return false, 0, 0
	}
}

func getCoordinatesForMoveRight(right_spaces int, x int, y int) (bool, int, int) {
	switch current_orientation {
	case NORTH:
		current_orientation = EAST
		return updateX(visited, x, x+right_spaces, y)
	case SOUTH:
		current_orientation = WEST
		return updateX(visited, x, x-right_spaces, y)
	case EAST:
		current_orientation = SOUTH
		return updateY(visited, x, y, y-right_spaces)
	case WEST:
		current_orientation = NORTH
		return updateY(visited, x, y, y+right_spaces)
	default:
		return false, 0, 0
	}
}

func updateX(locations map[string]bool, x int, new_x int, y int) (bool, int, int) {
	var start, end int = 0, 0
	if x < new_x {
		start = x + 1
		end = new_x
		for i := start; i <= end; i++ {
			location := strconv.Itoa(i) + strconv.Itoa(y)
			if visited[location] {
				return true, i, y
			} else {
				visited[location] = true
			}
		}
	} else {
		start = x - 1
		end = new_x
		for i := start; i >= end; i-- {
			location := strconv.Itoa(i) + strconv.Itoa(y)
			if visited[location] {
				return true, i, y
			} else {
				visited[location] = true
			}
		}
	}
	return false, new_x, y
}

func updateY(locations map[string]bool, x int, y int, new_y int) (bool, int, int) {
	var start, end int = 0, 0
	if y < new_y {
		start = y + 1
		end = new_y
		for i := start; i <= end; i++ {
			location := strconv.Itoa(x) + strconv.Itoa(i)
			if visited[location] {
				return true, x, i
			} else {
				visited[location] = true
			}
		}
	} else {
		start = y - 1
		end = new_y
		for i := start; i >= end; i-- {
			location := strconv.Itoa(x) + strconv.Itoa(i)
			if visited[location] {
				return true, x, i
			} else {
				visited[location] = true
			}
		}
	}
	return false, x, new_y
}
