package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
)

func main() {
	lines := utils.MustReadLines("day3/input.txt")

	// Part 1
	total := processLines(lines, `mul\([0-9]+,[0-9]+\)`, true)
	println(total)

	// Part 2
	total = processLines(lines, `(mul\([0-9]+,[0-9]+\))|(do\(\))|(don\'t\(\))`, false)
	fmt.Println(total)
}

func processLines(lines []string, pattern string, alwaysProcess bool) int {
	r, _ := regexp.Compile(pattern)
	total := 0
	isDo := true

	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, match := range matches {
			if match == "do()" {
				isDo = true
			} else if match == "don't()" {
				isDo = false
			} else if alwaysProcess || isDo {
				total += processMatch(match)
			}
		}
	}
	return total
}

func processMatch(match string) int {
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	s1, s2 := r.FindStringSubmatch(match)[1], r.FindStringSubmatch(match)[2]
	return utils.MustParseInt(s1) * utils.MustParseInt(s2)
}
