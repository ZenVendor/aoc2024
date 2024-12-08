package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Antenna struct {
	freq rune
	x    int
	y    int
}

type Antennae []Antenna

func (ants Antennae) Sort() {
	slices.SortFunc(ants, func(a, b Antenna) int {
		if n := int(a.freq) - int(b.freq); n != 0 {
			return n
		}
		if n := a.y - b.y; n != 0 {
			return n
		}
		return a.x - b.x
	})
}

func (ants Antennae) CheckLocation(x, y int) int {
	return slices.IndexFunc(ants, func(a Antenna) bool {
		return a.x == x && a.y == y
	})
}

func (ants Antennae) FindFreq(freq rune) int {
	return slices.IndexFunc(ants, func(a Antenna) bool {
		return a.freq == freq
	})
}

func (ants Antennae) FindAllOnFreq(freq rune) Antennae {
	ants.Sort()
	var rs Antennae
	idx := ants.FindFreq(freq)
	for i := idx; i < len(ants); i++ {
		if ants[i].freq == freq {
			rs = append(rs, ants[i])
		}
	}
	return rs
}

func (a Antenna) WithinField(w, h int) bool {
	if a.x < 0 || a.x >= w {
		return false
	}
	if a.y < 0 || a.y >= h {
		return false
	}
	return true
}

func day08(part int, file *os.File) {

	var ants Antennae
	var frequencies []rune
	fieldW := 0
	fieldH := 0

	x, y := 0, 0
	s := bufio.NewReader(file)
	for {
		char, _, err := s.ReadRune()
		if err != nil {
			break
		}
		if char == '\n' {
			fieldW = x
			fieldH = y + 1
			x = 0
			y++
			continue
		}
		if char != '.' {
			ants = append(ants, Antenna{freq: char, x: x, y: y})
			if !slices.Contains(frequencies, char) {
				frequencies = append(frequencies, char)
			}
		}
		x++
	}

	count := 0
	var antinodes Antennae
	if part == 1 {
		for _, freq := range frequencies {
			freqAnts := ants.FindAllOnFreq(freq)
			freqANs := Antennae{}
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
						if an.WithinField(fieldW, fieldH) {
							freqANs = append(freqANs, an)
						}

						an = freqAnts[j]
						an.x += xDist
						an.y += yDist
						if an.WithinField(fieldW, fieldH) {
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
			freqAnts := ants.FindAllOnFreq(freq)
			freqANs := Antennae{}
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
							if !an.WithinField(fieldW, fieldH) {
								break
							}
							freqANs = append(freqANs, an)
						}

						an = freqAnts[i]
						for {
							an.x = an.x + xDist
							an.y = an.y + yDist
							if !an.WithinField(fieldW, fieldH) {
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
	antinodes.Sort()
	var validANs Antennae
	for _, an := range antinodes {
		if validANs.CheckLocation(an.x, an.y) == -1 {
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
				if idx := ants.CheckLocation(x, y); idx != -1 {
					char = string(ants[idx].freq)
				}
				fmt.Printf("%s", char)
			}
			fmt.Printf("\n")
		}
	*/
	fmt.Printf("Day 08 part %d: %d\n", part, count)

}
