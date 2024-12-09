package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func TryOperations(values []int, maxIdx, idx, total, testVal int, ops []string, part int) ([]string, int, bool) {
	currVal := values[idx]

	currTotal := total * currVal
	ops = append(ops, "*")

	if idx == maxIdx && currTotal == testVal {
		return ops, currTotal, true
	}
	if idx < maxIdx {
		ops, currTotal, ok := TryOperations(values, maxIdx, idx+1, currTotal, testVal, ops, part)
		if ok {
			return ops, currTotal, true
		}
	}
	ops = ops[:len(ops)-1]

	currTotal = total + currVal
	ops = append(ops, "+")

	if idx == maxIdx && currTotal == testVal {
		return ops, currTotal, true
	}
	if idx < maxIdx {
		ops, currTotal, ok := TryOperations(values, maxIdx, idx+1, currTotal, testVal, ops, part)
		if ok {
			return ops, currTotal, true
		}
	}
	ops = ops[:len(ops)-1]

	if part == 2 {
		currTotal, _ = strconv.Atoi(fmt.Sprintf("%d%d", total, currVal))
		ops = append(ops, "||")

		if idx == maxIdx && currTotal == testVal {
			return ops, currTotal, true
		}
		if idx < maxIdx {
			ops, currTotal, ok := TryOperations(values, maxIdx, idx+1, currTotal, testVal, ops, part)
			if ok {
				return ops, currTotal, true
			}
		}
		ops = ops[:len(ops)-1]
	}

	return ops, total, false
}

func day07(part int, file *os.File) {

	sum := 0

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := strings.Split(s.Text(), ": ")
		tv, err := strconv.Atoi(line[0])
		if err != nil {
			log.Fatal(err)
		}
		vals := []int{}
		for _, n := range strings.Split(line[1], " ") {
			val, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, val)
		}

		ops := []string{}
		ops, _, ok := TryOperations(vals, len(vals)-1, 1, vals[0], tv, ops, part)
		if ok {
			sum += tv
		}
	}

	fmt.Printf("Day 07 part %d: %d\n", part, sum)
}
