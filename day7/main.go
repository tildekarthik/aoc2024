package main

import (
	"aoc2024/utils"
	"strconv"
	"strings"
)

func main() {
	input := utils.MustReadLines("day7/input.txt")
	ops := []string{"+", "*", "C"}
	rsum := 0
	for _, line := range input {
		rsum += getResulIfValid(line, ops)
	}
	println("Sum of valid lines: ", rsum)
}

func getResulIfValid(line string, ops []string) int {
	nums := getNums(line)
	result := getResult(line)
	opCombs := createComb(ops, len(nums)-1)
	for _, opComb := range opCombs {
		if evalExpr(nums, opComb) == result {
			println("Found the correct combination: ", opComb, "for line", line)
			return result
		}
	}
	return 0
}

func createComb(ops []string, posns int) []string {
	combs := []string{}
	combs = append(combs, ops...)
	for i := 0; i < posns-1; i++ {
		var c []string
		for _, op := range ops {
			for _, comb := range combs {
				c = append(c, comb+op)
			}
		}
		combs = c
	}
	return combs
}

func evalExpr(nums []int, opsS string) int {
	ops := []string{}

	for _, op := range opsS {
		ops = append(ops, string(op))
	}

	result := nums[0]
	for i, num := range nums[1:] {
		switch ops[i] {
		case "+":
			result += num
		case "*":
			result *= num
		case "C":
			// concatenates result with num
			result = utils.MustParseInt(strconv.Itoa(result) + strconv.Itoa(num))
		}
	}
	return result
}

func getResult(line string) int {
	// split the line initially using ": "
	lh := strings.Split(line, ": ")[0]
	result := utils.MustParseInt(lh)
	return result
}

func getNums(line string) []int {
	r := strings.Split(line, ": ")[1]
	nums := []int{}
	for _, n := range strings.Split(r, " ") {
		nums = append(nums, utils.MustParseInt(n))
	}
	return nums
}
