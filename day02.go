package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func RuleCheck(x, y int) (dir int, ok bool) {
	if x == y {
		return 0, false
	}
	ok = true
	if x > y {
		dir = -1
		if x-y > 3 {
			ok = false
		}
	}
	if x < y {
		dir = 1
		if y-x > 3 {
			ok = false
		}
	}
	return dir, ok
}

func RecheckReport(report []int, ignore, main_dir int) bool {
	var cp []int
	for i, num := range report {
		if i == ignore {
			continue
		}
		cp = append(cp, num)
	}

	for i := 0; i < len(cp)-1; i++ {
		dir, ok := RuleCheck(cp[i], cp[i+1])
		if !ok || dir != main_dir {
			return false
		}
	}
	return true
}

func CheckReport(report []int) (int, bool) {
	var dirs []int
	var bad []int
	count_inc := 0
	count_dec := 0
	count_nc := 0
	count_err := 0
	for i := 0; i < len(report)-1; i++ {
		curr := report[i]
		next := report[i+1]
		dir, ok := RuleCheck(curr, next)
		dirs = append(dirs, dir)
		switch dir {
		case -1:
			count_dec++
		case 0:
			count_nc++
		case 1:
			count_inc++
		}
		if !ok {
			count_err++
			if !slices.Contains(bad, i) {
				bad = append(bad, i)
			}
			if !slices.Contains(bad, i+1) {
				bad = append(bad, i+1)
			}
		}
	}
	if len(bad) == 0 && count_inc*count_dec+count_nc == 0 {
		return count_err, true
	}
	if count_inc == 1 && count_nc == 0 {
		idx := slices.Index(dirs, 1)
		if !slices.Contains(bad, idx) {
			bad = append(bad, idx)
		}
		if !slices.Contains(bad, idx+1) {
			bad = append(bad, idx+1)
		}
	}
	if count_dec == 1 && count_nc == 0 {
		idx := slices.Index(dirs, -1)
		if !slices.Contains(bad, idx) {
			bad = append(bad, idx)
		}
		if !slices.Contains(bad, idx+1) {
			bad = append(bad, idx+1)
		}
	}
	main_dir := 1
	if count_inc-count_dec < 0 {
		main_dir = -1
	}
	if len(bad) > 0 && count_nc < 2 && (count_inc < 2 || count_dec < 2) {
		for _, idx := range bad {
			ok := RecheckReport(report, idx, main_dir)
			if ok {
				for i, r := range report {
					if i == idx {
						fmt.Printf("\033[1m%d\033[22m ", r)
					} else {
						fmt.Printf("%d ", r)
					}
				}
				fmt.Printf(" :: ignored, %d\n", count_err)
				return count_err, true
			}
		}
	}
	return count_err, false
}

func day02(part int, file *os.File) {

	safe := 0

	s := bufio.NewScanner(file)
	for s.Scan() {
		nums := strings.Split(s.Text(), " ")
		var report []int
		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			report = append(report, n)
		}
		errs, ok := CheckReport(report)
		if ok {
			safe++
			fmt.Printf("%v\t%v, %t, %d\n", nums, report, ok, errs)
		}
	}
	fmt.Printf("%d\n", safe)
}
