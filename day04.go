package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day04(part int, file *os.File) {

	lines := []string{}
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if part == 1 {

		allLines := []string{}
		for _, l := range lines {
			allLines = append(allLines, l)
		}

		for x := 0; x < len(lines[0]); x++ {
			var temp string
			for y := 0; y < len(lines); y++ {
				temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
			}
			allLines = append(allLines, temp)
		}

		for i := 0; i < len(lines[0]); i++ {
			var temp string
			x, y := i, 0
			for {
				if x == len(lines[0]) || y == len(lines) {
					break
				}
				temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
				x++
				y++
			}
			if len(temp) >= 4 {
				allLines = append(allLines, temp)
			}
		}
		for i := 1; i < len(lines); i++ {
			var temp string
			x, y := 0, i
			for {
				if x == len(lines[0]) || y == len(lines) {
					break
				}
				temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
				x++
				y++
			}
			if len(temp) >= 4 {
				allLines = append(allLines, temp)
			}
		}

		for i := len(lines[0]) - 1; i >= 0; i-- {
			var temp string
			x, y := i, 0
			for {
				if x < 0 || y == len(lines) {
					break
				}
				temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
				x--
				y++
			}
			if len(temp) >= 4 {
				allLines = append(allLines, temp)
			}
		}
		for i := 1; i < len(lines); i++ {
			var temp string
			x, y := len(lines)-1, i
			for {
				if x < 0 || y == len(lines) {
					break
				}
				temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
				x--
				y++
			}
			if len(temp) >= 4 {
				allLines = append(allLines, temp)
			}
		}
		for _, l := range allLines {
			str := strings.Count(l, "XMAS")
			rev := strings.Count(l, "SAMX")
			count += str + rev
		}
	}

	if part == 2 {
		for y := 1; y < len(lines)-1; y++ {
			for x := 1; x < len(lines[0])-1; x++ {
				if lines[y][x] == 'A' {
					xmas := fmt.Sprintf("%sA%s%sA%s",
						string(lines[y-1][x-1]),
						string(lines[y+1][x+1]),
						string(lines[y-1][x+1]),
						string(lines[y+1][x-1]),
					)
					str := strings.Count(xmas, "MAS")
					rev := strings.Count(xmas, "SAM")
					if str+rev == 2 {
						count++
					}
				}
			}
		}
	}
	fmt.Printf("Day 04 part %d: %d\n", part, count)
}
