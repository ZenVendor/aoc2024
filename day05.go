package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	page  int
	after []int
}

type Rules []Rule

func (rules Rules) PageInRules(page int) int {
	return slices.IndexFunc(rules, func(rule Rule) bool {
		if rule.page == page {
			return true
		}
		return false
	})
}

func (rules Rules) UpdateRule(first, second int) []Rule {
	if idx := rules.PageInRules(first); idx == -1 {
		rules = append(rules, Rule{page: first, after: []int{second}})
	} else {
		if !slices.Contains(rules[idx].after, second) {
			rules[idx].after = append(rules[idx].after, second)
		}
	}
	return rules
}

type Update struct {
	pages     []int
	isCorrect bool
	incorrect Positions
}

func (u Update) Check(rules Rules) Update {
	for i := 0; i < len(u.pages); i++ {
		idx := rules.PageInRules(u.pages[i])
		if idx == -1 {
			continue
		}
		rule := rules[idx]
		for j := 0; j < i; j++ {
			if slices.Contains(rule.after, u.pages[j]) {
				p := Position{x: u.pages[i], y: u.pages[j], value: 0}
				if !u.incorrect.ContainsPair(p) {
					u.incorrect = append(u.incorrect, p)
				}
			}
		}
	}
	if len(u.incorrect) > 0 {
		u.isCorrect = false
		return u
	}
	u.isCorrect = true
	return u
}

func (u Update) Fix() Update {
	for _, pair := range u.incorrect {
		first := slices.Index(u.pages, pair.x)
		second := slices.Index(u.pages, pair.y)
		u.pages[first], u.pages[second] = u.pages[second], u.pages[first]
	}
	u.incorrect = Positions{}
	return u
}

func day05(part int, file *os.File) {

	var rules Rules
	var updates []Update
	section := 1

	s := bufio.NewScanner(file)
	for s.Scan() {
		if s.Text() == "" {
			section = 2
			continue
		}
		if section == 1 {
			pair := strings.Split(s.Text(), "|")
			first, err := strconv.Atoi(pair[0])
			if err != nil {
				log.Fatal(err)
			}
			second, err := strconv.Atoi(pair[1])
			if err != nil {
				log.Fatal(err)
			}
			rules = rules.UpdateRule(first, second)
		}
		if section == 2 {
			pages := []int{}
			line := strings.Split(s.Text(), ",")

			for _, num := range line {
				num, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				pages = append(pages, num)
			}
			updates = append(updates, Update{pages: pages})
		}
	}

	midpages := 0

	if part == 1 {
		for uid := 0; uid < len(updates); uid++ {
			updates[uid] = updates[uid].Check(rules)
			if updates[uid].isCorrect {
				mid := len(updates[uid].pages) / 2
				midpages += updates[uid].pages[mid]
			}
		}
	}

	if part == 2 {
		for uid := 0; uid < len(updates); uid++ {
			updates[uid] = updates[uid].Check(rules)
			if updates[uid].isCorrect {
				continue
			}
			for {
				updates[uid] = updates[uid].Check(rules)
				if updates[uid].isCorrect {
					break
				}
				updates[uid] = updates[uid].Fix()
			}
			mid := len(updates[uid].pages) / 2
			midpages += updates[uid].pages[mid]
		}
	}
	fmt.Printf("Day 05 part %d: %d\n", part, midpages)
}
