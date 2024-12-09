package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func day08(part int, file *os.File) {

	var field Plan
	var frequencies []int

	x, y := 0, 0
	s := bufio.NewReader(file)
	for {
		char, _, err := s.ReadRune()
		if err != nil {
			break
		}
		if char == '\n' {
			field.width = x
			field.height = y + 1
			x = 0
			y++
			continue
		}
		if char != '.' {
			val := int(char)
			field.objects = append(field.objects, Position{x: x, y: y, value: val})
			if !slices.Contains(frequencies, val) {
				frequencies = append(frequencies, val)
			}
		}
		x++
	}

	count := 0
	var antinodes Positions
	if part == 1 {
		for _, freq := range frequencies {
			freqAnts := field.objects.FindAllWithValue(freq)
			var freqANs Positions
			if len(freqAnts) > 1 {
				for i := 0; i < len(freqAnts); i++ {
					for j := 0; j < len(freqAnts); j++ {
						if i == j {
							continue
						}
						xDist := freqAnts[j].x - freqAnts[i].x
						yDist := freqAnts[j].y - freqAnts[i].y

						an := freqAnts[i]
						an.x -= xDist
						an.y -= yDist
						if !an.OutOfBounds(field.width, field.height) {
							freqANs = append(freqANs, an)
						}

						an = freqAnts[j]
						an.x += xDist
						an.y += yDist
						if !an.OutOfBounds(field.width, field.height) {
							freqANs = append(freqANs, an)
						}
					}
				}
			}
			antinodes = slices.Concat(antinodes, freqANs)
		}
	}
	if part == 2 {
		for _, freq := range frequencies {
			freqAnts := field.objects.FindAllWithValue(freq)
			var freqANs Positions
			if len(freqAnts) > 1 {
				for i := 0; i < len(freqAnts); i++ {
					for j := 0; j < len(freqAnts); j++ {
						if i == j {
							continue
						}
						xDist := freqAnts[j].x - freqAnts[i].x
						yDist := freqAnts[j].y - freqAnts[i].y

						an := freqAnts[j]
						for {
							an.x = an.x - xDist
							an.y = an.y - yDist
							if an.OutOfBounds(field.width, field.height) {
								break
							}
							freqANs = append(freqANs, an)
						}

						an = freqAnts[i]
						for {
							an.x = an.x + xDist
							an.y = an.y + yDist
							if an.OutOfBounds(field.width, field.height) {
								break
							}
							freqANs = append(freqANs, an)
						}
					}
				}
			}
			antinodes = slices.Concat(antinodes, freqANs)
		}
	}
	antinodes.SortByValueYX()
	var validANs Positions
	for _, an := range antinodes {
		if !validANs.Contains(an) {
			validANs = append(validANs, an)
			count++
		}
	}
	/*
		for y := 0; y < fieldH; y++ {
			for x := 0; x < fieldW; x++ {
				char := "."
				if idx := validANs.CheckLocation(x, y); idx != -1 {
					char = "#"
				}
				if idx := field.objects.CheckLocation(x, y); idx != -1 {
					char = string(field.objects[idx].freq)
				}
				fmt.Printf("%s", char)
			}
			fmt.Printf("\n")
		}
	*/
	fmt.Printf("Day 08 part %d: %d\n", part, count)

}
