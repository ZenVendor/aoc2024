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
	page int
	post []int
}

func PageInRules(rules []Rule, page int) int {
	return slices.IndexFunc(rules, func(rule Rule) bool {
		if rule.page == page {
			return true
		}
		return false
	})
}

func UpdateRule(rules []Rule, first, second int) []Rule {
	if idx := PageInRules(rules, first); idx == -1 {
		rules = append(rules, Rule{page: first, post: []int{second}})
	} else {
		if !slices.Contains(rules[idx].post, second) {
			rules[idx].post = append(rules[idx].post, second)
		}
	}
	return rules
}

type Update struct {
	pages     []int
	isCorrect bool
	incorrect []Position
}

func PairInList(list []Position, pair Position) bool {
	return slices.ContainsFunc(list, func(p Position) bool {
		return (p.x == pair.x && p.y == pair.y) ||
			(p.x == pair.y && p.y == pair.x)
	})
}

func Check(u Update, rules []Rule) Update {
	for i := 0; i < len(u.pages); i++ {
		idx := PageInRules(rules, u.pages[i])
		if idx == -1 {
			continue
		}
		rule := rules[idx]
		for j := 0; j < i; j++ {
			if slices.Contains(rule.post, u.pages[j]) {
				p := Position{u.pages[i], u.pages[j]}
				if !PairInList(u.incorrect, p) {
					u.incorrect = append(u.incorrect, p)
					//fmt.Printf("\tIncorrect: %d, %d\n", p.first, p.second)
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

func Fix(u Update) Update {
	for _, pair := range u.incorrect {
		first := slices.Index(u.pages, pair.x)
		second := slices.Index(u.pages, pair.y)
		u.pages[first], u.pages[second] = u.pages[second], u.pages[first]
	}
	u.incorrect = []Position{}
	return u
}

func day05(part int, file *os.File) {

	rules := []Rule{}
	updates := []Update{}
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
			rules = UpdateRule(rules, first, second)
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
			updates[uid] = Check(updates[uid], rules)
			if updates[uid].isCorrect {
				mid := len(updates[uid].pages) / 2
				midpages += updates[uid].pages[mid]
			}
		}
	}

	if part == 2 {
		for uid := 0; uid < len(updates); uid++ {
			updates[uid] = Check(updates[uid], rules)
			if updates[uid].isCorrect {
				continue
			}
			for {
				updates[uid] = Check(updates[uid], rules)
				if updates[uid].isCorrect {
					break
				}
				updates[uid] = Fix(updates[uid])
			}
			mid := len(updates[uid].pages) / 2
			midpages += updates[uid].pages[mid]
		}
	}
	fmt.Printf("Day 05 part %d: %d\n", part, midpages)
}
