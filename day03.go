package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day03(part int, file *os.File) {

	enabled := true
	total := 0

	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString(')')
		if err != nil {
			break
		}
		fmt.Printf("%s\n", line)
		if len(line) < 4 {
			fmt.Println("\033[1mToo short\033[22m")
			continue
		}
		if len(line) > 6 && strings.Compare(line[len(line)-7:], "don't()") == 0 {
			enabled = false
			fmt.Printf("\033[1mDisabling\033[22m\n")
			continue
		}
		if len(line) > 3 && strings.Compare(line[len(line)-4:], "do()") == 0 {
			enabled = true
			fmt.Printf("\033[1mEnabling\033[22m\n")
			continue
		}
		lo := 0
		hi := len(line) - 1
		for {
			part := line[lo:hi]
			idx := strings.Index(part, "mul")
			if idx == -1 {
				break
			}
			if part[idx+3] == '(' {
				nums := strings.Split(part[idx+4:], ",")
				if len(nums) != 2 {
					lo += idx + 3
					continue
				}
				x, err := strconv.Atoi(nums[0])
				if err != nil {
					lo += idx + 3
					continue
				}
				y, err := strconv.Atoi(nums[1])
				if err != nil {
					lo += idx + 3
					continue
				}
				if enabled {
					total += x * y
				}
				fmt.Printf("\033[1mx: %d, y: %d, enabled: %t\033[22m\n", x, y, enabled)

			}
			lo += idx + 3
		}
	}
	fmt.Printf("Total: %d", total)

}
