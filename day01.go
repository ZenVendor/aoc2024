package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day01(part int, file *os.File) {

	var ids1 []int
	var ids2 []int

	s := bufio.NewScanner(file)
	for s.Scan() {
		nums := strings.Split(s.Text(), "   ")
		id1, _ := strconv.Atoi(nums[0])
		id2, _ := strconv.Atoi(nums[1])
		ids1 = append(ids1, id1)
		ids2 = append(ids2, id2)
	}
	slices.Sort(ids1)
	slices.Sort(ids2)
	result := 0
	for n := 0; n < len(ids1); n++ {
		diff := ids1[n] - ids2[n]
		if diff < 0 {
			diff = -diff
		}
		result += diff
	}

	fmt.Printf("Day 01 result 1: %d\n", result)

	if part == 2 {
		score := 0
		for _, id := range ids1 {
			count := 0
			if idx := slices.Index(ids2, id); idx != -1 {
				for i := idx; i < len(ids2); i++ {
					if ids2[i] != id {
						break
					}
					count++
				}
			}
			score += id * count
		}
		fmt.Printf("Day 01 result 2: %d\n", score)
	}

}
