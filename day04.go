package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func day04(part int, file *os.File) {

	count := 0
	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		str := strings.Count(scanner.Text(), "XMAS")
		rev := strings.Count(scanner.Text(), "SAMX")
		count += str + rev
		//fmt.Printf("%s :: str: %d, rev: %d\n", scanner.Text(), str, rev)
	}

	for x := 0; x < len(lines[0]); x++ {
		var temp string
		for y := 0; y < len(lines); y++ {
			temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
		}

		str := strings.Count(temp, "XMAS")
		rev := strings.Count(temp, "SAMX")
		count += str + rev
		//fmt.Printf("%s :: str: %d, rev: %d\n", temp, str, rev)
	}

	for i := 0; i < len(lines[0]); i++ {
		var temp string
		x := i
		y := 0
		for {
			if x == len(lines[0]) || y == len(lines) {
				break
			}
			temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
			x++
			y++
		}
		if len(temp) < 4 {
			continue
		}
		str := strings.Count(temp, "XMAS")
		rev := strings.Count(temp, "SAMX")
		count += str + rev
		//fmt.Printf("%s :: str: %d, rev: %d\n", temp, str, rev)
	}
	for i := 1; i < len(lines); i++ {
		var temp string
		x := 0
		y := i
		for {
			if x == len(lines[0]) || y == len(lines) {
				break
			}
			temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
			x++
			y++
		}
		if len(temp) < 4 {
			continue
		}
		str := strings.Count(temp, "XMAS")
		rev := strings.Count(temp, "SAMX")
		count += str + rev
		//fmt.Printf("%s :: str: %d, rev: %d\n", temp, str, rev)
	}

	for i := len(lines[0]) - 1; i >= 0; i-- {
		var temp string
		x := i
		y := 0
		for {
			if x < 0 || y == len(lines) {
				break
			}
			temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
			x--
			y++
		}
		if len(temp) < 4 {
			continue
		}
		str := strings.Count(temp, "XMAS")
		rev := strings.Count(temp, "SAMX")
		count += str + rev
		//fmt.Printf("%s :: str: %d, rev: %d\n", temp, str, rev)
	}
	for i := 1; i < len(lines); i++ {
		var temp string
		x := len(lines) - 1
		y := i
		for {
			if x < 0 || y == len(lines) {
				break
			}
			temp = fmt.Sprintf("%s%s", temp, string(lines[y][x]))
			x--
			y++
		}
		if len(temp) < 4 {
			continue
		}
		str := strings.Count(temp, "XMAS")
		rev := strings.Count(temp, "SAMX")
		count += str + rev
		//fmt.Printf("%s :: str: %d, rev: %d\n", temp, str, rev)
	}
	fmt.Printf("Day 04 part 1: %d\n", count)

	if part == 2 {
		xmascount := 0
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
						xmascount++
					}
				}
			}
		}
		fmt.Printf("Day 04 part 2: %d\n", xmascount)
	}
}
