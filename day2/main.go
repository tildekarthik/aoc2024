package main

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func main() {
	fname := "day2/input.txt"
	reports := utils.MustReadLines(fname)
	safeCnt := 0
	for _, report := range reports {
		reportNums := parseLine(report)
		if isSafe(reportNums) {
			safeCnt++
		}
	}
	fmt.Println("Safe reports:", safeCnt)
}

func parseLine(report string) []int {
	reportSplit := strings.Split(report, " ")
	var reportNums []int
	for _, num := range reportSplit {
		reportNums = append(reportNums, utils.MustParseInt(num))
	}
	return reportNums
}

func diffLevels(levels []int) []int {
	var diffLevels []int
	if len(levels) == 1 {
		return diffLevels
	}
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		diffLevels = append(diffLevels, diff)
	}
	return diffLevels
}

func testIfSafe(d []int) bool {
	if len(d) == 0 {
		return true
	}
	isPositive := d[0] > 0
	for _, num := range d {
		if (isPositive && num <= 0) || (!isPositive && num >= 0) || (isPositive && num > 3) || (!isPositive && num < -3) {
			return false
		}
	}
	return true
}

func checkByDroppingOneElement(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		copyLevels := append(levels[:i], levels[i+1:]...)
		if testIfSafe(diffLevels(copyLevels)) {
			return true
		}
	}
	return false
}

func isSafe(reportNums []int) bool {
	if testIfSafe(diffLevels(reportNums)) {
		return true
	}
	return checkByDroppingOneElement(reportNums)
}
