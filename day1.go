package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*
--- Day 1: No Time for a Taxicab ---

Santa's sleigh uses a very high-precision clock to guide its movements, and the clock's oscillator is regulated by stars. Unfortunately, the stars have been stolen... by the Easter Bunny. To save Christmas, Santa needs you to retrieve all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You're airdropped near Easter Bunny Headquarters in a city somewhere. "Near", unfortunately, is as close as you can get - the instructions on the Easter Bunny Recruiting Document the Elves intercepted start here, and nobody had time to work them out further.

The Document indicates that you should start at the given coordinates (where you just landed) and face North. Then, follow the provided sequence: either turn left (L) or right (R) 90 degrees, then walk forward the given number of blocks, ending at a new intersection.

There's no time to follow such ridiculous instructions on foot, though, so you take a moment and work out the destination. Given that you can only walk on the street grid of the city, how far is the shortest path to the destination?

For example:

Following R2, L3 leaves you 2 blocks East and 3 blocks North, or 5 blocks away.
R2, R2, R2 leaves you 2 blocks due South of your starting position, which is 2 blocks away.
R5, L5, R5, R3 leaves you 12 blocks away.
How many blocks away is Easter Bunny HQ?

To begin, get your puzzle input.

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

func main() {
	input := "R4, R1, L2, R1, L1, L1, R1, L5, R1, R5, L2, R3, L3, L4, R4, R4, R3, L5, L1, R5, R3, L4, R1, R5, L1, R3, L2, R3, R1, L4, L1, R1, L1, L5, R1, L2, R2, L3, L5, R1, R5, L1, R188, L3, R2, R52, R5, L3, R79, L1, R5, R186, R2, R1, L3, L5, L2, R2, R4, R5, R5, L5, L4, R5, R3, L4, R4, L4, L4, R5, L4, L3, L1, L4, R1, R2, L5, R3, L4, R3, L3, L5, R1, R1, L3, R2, R1, R2, R2, L4, R5, R1, R3, R2, L2, L2, L1, R2, L1, L3, R5, R1, R4, R5, R2, R2, R4, R4, R1, L3, R4, L2, R2, R1, R3, L5, R5, R2, R5, L1, R2, R4, L1, R5, L3, L3, R1, L4, R2, L2, R1, L1, R4, R3, L2, L3, R3, L2, R1, L4, R5, L1, R5, L2, L1, L5, L2, L5, L2, L4, L2, R3"
	distance := distanceToBunnyHeadquarters(input)
	fmt.Printf("The Distance to Bunny HQ is %d\n", distance)
}

func distanceToBunnyHeadquarters(input string) int {
	var curr_x, curr_y int = 0, 0
	coordinates := strings.Split(input, ", ")
	for i := 0; i < len(coordinates); i++ {
		dir, spaces := getSpecs(coordinates[i])
		if dir == "L" {
			curr_x, curr_y = getCoordinatesForMoveLeft(spaces, curr_x, curr_y)
		} else {
			curr_x, curr_y = getCoordinatesForMoveRight(spaces, curr_x, curr_y)
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

func getCoordinatesForMoveLeft(left_spaces int, x int, y int) (int, int) {
	switch current_orientation {
	case NORTH:
		current_orientation = WEST
		return x - left_spaces, y
	case SOUTH:
		current_orientation = EAST
		return x + left_spaces, y
	case EAST:
		current_orientation = NORTH
		return x, y + left_spaces
	case WEST:
		current_orientation = SOUTH
		return x, y - left_spaces
	default:
		return 0, 0
	}
}

func getCoordinatesForMoveRight(right_spaces int, x int, y int) (int, int) {
	switch current_orientation {
	case NORTH:
		current_orientation = EAST
		return x + right_spaces, y
	case SOUTH:
		current_orientation = WEST
		return x - right_spaces, y
	case EAST:
		current_orientation = SOUTH
		return x, y - right_spaces
	case WEST:
		current_orientation = NORTH
		return x, y + right_spaces
	default:
		return 0, 0
	}
}
