package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func PositionIndex(obs []Position, pos Position) int {
	return slices.IndexFunc(obs, func(p Position) bool {
		return p.x == pos.x && p.y == pos.y
	})
}

func day06(part int, file *os.File) {

	var plan [][]rune
	var obstacles []Position
	var currPos Position
	var currDir int

	row := 0
	col := 0
	var line []rune

	reader := bufio.NewReader(file)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if char == '\n' {
			plan = append(plan, line)
			row++
			col = 0
			line = []rune{}
			continue
		}
		line = append(line, char)
		if char == '#' {
			obstacles = append(obstacles, Position{col, row})
		}
		if char == '^' {
			currPos = Position{col, row}
			currDir = UP
		}
		if char == '>' {
			currPos = Position{col, row}
			currDir = RIGHT
		}
		if char == 'v' {
			currPos = Position{col, row}
			currDir = DOWN
		}
		if char == '<' {
			currPos = Position{col, row}
			currDir = LEFT
		}
		col++
	}

	var hLines []Position
	var vLines []Position
	var path []Position
	var newObs []Position

	for {
		if PositionIndex(path, currPos) == -1 {
			path = append(path, currPos)
		}
		if (currDir == UP || currDir == DOWN) && !slices.Contains(vLines, Position{currPos.x, currDir}) {
			vLines = append(vLines, Position{currPos.x, currDir})
		}
		if (currDir == RIGHT || currDir == LEFT) && !slices.Contains(hLines, Position{currPos.y, currDir}) {
			hLines = append(hLines, Position{currPos.y, currDir})
		}

		nextPos := Position{}
		nextPos.x = currPos.x + translateDir()[currDir].x
		nextPos.y = currPos.y + translateDir()[currDir].y

		if nextPos.x < 0 || nextPos.x >= len(plan[0]) {
			break
		}
		if nextPos.y < 0 || nextPos.y >= len(plan) {
			break
		}
		if PositionIndex(obstacles, nextPos) > -1 {
			currDir = (currDir + 1) % 4
			continue
		}

		if PositionIndex(vLines, Position{currPos.x, (currDir + 1) % 4}) > -1 {
			if PositionIndex(newObs, nextPos) == -1 {
				newObs = append(newObs, nextPos)
			}
		}
		if PositionIndex(hLines, Position{currPos.y, (currDir + 1) % 4}) > -1 {
			if PositionIndex(newObs, nextPos) == -1 {
				newObs = append(newObs, nextPos)
			}
		}
		currPos = nextPos
	}

	fmt.Printf("Day 06 part 1: %d\n", len(path))

	if part == 2 {
		fmt.Printf("Day 06 part 2: %d\n", len(newObs))
	}

	//fmt.Printf("\x1bc")
	for y, row := range plan {
		for x, col := range row {
			char := string(col)

			if PositionIndex(path, Position{x, y}) > -1 {
				char = "X"
			}
			if part == 2 {
				if PositionIndex(newObs, Position{x, y}) > -1 {
					char = "O"
				}
			}
			if x == currPos.x && y == currPos.y {
				char = displayDir()[currDir]
			}

			fmt.Printf("%s", char)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%v\n", vLines)
	fmt.Printf("%v\n", hLines)

}
