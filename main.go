package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var file *os.File

func PrintHelp() {
	fmt.Printf("Usage: aoc2024 [day] [part] [file]\n")
	fmt.Printf("\tday: int 1-25\n")
	fmt.Printf("\tpart: int 1-2\n")
	fmt.Printf("\tfile string \"sample\"|\"data\"\n")
}

func main() {
	if len(os.Args) < 3 || len(os.Args) > 4 {
		PrintHelp()
		return
	}

	args := os.Args[1:]
	day, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
	}
	if day < 1 || day > 25 {
		PrintHelp()
		return
	}

	part, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println(err)
	}
	if part < 1 || part > 2 {
		PrintHelp()
		return
	}

	if len(args) == 3 {
		filename := fmt.Sprintf("files/day%02d.%s", day, args[2])
		file, err = os.Open(filename)
		if err != nil {
			fmt.Printf("Cannot open file: %s\n", err)
			return
		}
		defer file.Close()
	}

	switch day {
	case 1:
		day01(part, file)
	case 2:
		day02(part, file)
	case 3:
		day03(part, file)
	case 4:
		day04(part, file)
	case 5:
		day05(part, file)
	case 6:
		day06(part, file)
	case 7:
		day07(part, file)
	case 8:
		day08(part, file)
	case 9:
		day09(part, file)
	case 10:
		day10(part, file)
	case 11:
		day11(part, file)
	case 12:
		day12(part, file)
	case 13:
		day13(part, file)
	case 14:
		day14(part, file)
	case 15:
		day15(part, file)
	case 16:
		day16(part, file)
	case 17:
		day17(part, file)
	case 18:
		day18(part, file)
	case 19:
		day19(part, file)
	case 20:
		day20(part, file)
	case 21:
		day21(part, file)
	case 22:
		day22(part, file)
	case 23:
		day23(part, file)
	case 24:
		day24(part, file)
	case 25:
		day25(part, file)
	}

	return
}
