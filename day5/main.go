package main

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strings"
)

type rule struct {
	l string
	r string
}

func main() {
	lines := utils.MustReadLines("day5/input.txt")
	split := slices.Index(lines, "")
	ruleStrings := lines[:split]
	updateStrings := lines[split+1:]
	rules := make([]rule, len(ruleStrings))
	for i, s := range ruleStrings {
		rules[i] = MustParseRule(s)
	}
	updates := make([][]string, len(updateStrings))
	for i, s := range updateStrings {
		updates[i] = MustParseUpdate(s)
	}

	count := 0
	countWrongUpdates := 0
	// var wrongUpdates [][]string
	for _, u := range updates {
		if isUpdateValid(u, rules) {
			count += getMiddlePageNumber(u)
		} else {
			countWrongUpdates += getMiddlePageNumber(getCorrectlyOrderedUpdate(u, rules))
		}

	}
	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", countWrongUpdates)
}

func MustParseRule(s string) rule {
	ss := strings.Split(s, "|")
	return rule{ss[0], ss[1]}
}

func MustParseUpdate(s string) []string {
	ss := strings.Split(s, ",")
	return ss
}

func isRuleApplicable(u []string, r rule) bool {
	return slices.Contains(u, r.l) && slices.Contains(u, r.r)
}

func isRuleValid(u []string, r rule) bool {
	lindex := slices.Index(u, r.l)
	rindex := slices.Index(u, r.r)
	return lindex < rindex
}

func isUpdateValid(u []string, rs []rule) bool {
	isValid := true
	for _, r := range rs {
		if isRuleApplicable(u, r) {
			if !isRuleValid(u, r) {
				isValid = false
				break
			}
		}
	}
	return isValid
}

func getMiddlePageNumber(u []string) int {
	middleIndex := len(u) / 2
	return utils.MustParseInt(u[middleIndex])
}

func swapElements(u []string, i int, j int) []string {
	u[i], u[j] = u[j], u[i]
	return u
}

func getCorrectlyOrderedUpdate(u []string, rs []rule) []string {
	for {
		for _, r := range rs {
			if isRuleApplicable(u, r) {
				if !isRuleValid(u, r) {
					lindex := slices.Index(u, r.l)
					rindex := slices.Index(u, r.r)
					u = swapElements(u, lindex, rindex)
					break
				}
			}
		}
		if isUpdateValid(u, rs) {
			break
		}
	}
	return u
}
