package main

import (
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fname := "day1/input.txt"
	lines := utils.MustReadLines(fname)
	left, right := parseLines(lines)

	sumDistances(left, right)
	similarityScore(left, right)
}

func parseLines(lines []string) ([]int, []int) {
	var left, right []int
	for _, line := range lines {
		l, r := parseLine(line)
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func sumDistances(left []int, right []int) {
	slices.Sort(left)
	slices.Sort(right)
	sumDist := 0
	for i, l := range left {
		sumDist += absInt(l - right[i])
	}
	fmt.Printf("Sum of distances: %d\n", sumDist)
}

func similarityScore(left []int, right []int) {
	similarity := 0
	for _, l := range left {
		similarity += l * countTimesInList(right, l)
	}
	fmt.Printf("Similarity: %d\n", similarity)
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func parseLine(line string) (int, int) {
	splits := strings.Split(line, "   ")
	return mustParseInt(splits[0]), mustParseInt(splits[1])
}

func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func countTimesInList(list []int, val int) int {
	count := 0
	for _, v := range list {
		if v == val {
			count++
		}
	}
	return count
}
