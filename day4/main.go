package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
)

type Point struct {
	x int
	y int
}

func main() {
	matrix := readIntoMatrix("day4/input.txt")
	var count int

	// Part 1
	count = 0
	Y := len(matrix)
	X := len(matrix[0])
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			p := Point{x, y}
			count += getCountsForLocation(matrix, p)
		}
	}
	// count = getCountsForLocation(matrix, Point{x: 4, y: 1})
	fmt.Println("Part 1: ", count)
	// Part 2
	count = 0
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			p := Point{x, y}
			if isXMAS(matrix, p) {
				count++
			}
		}
	}
	fmt.Println("Part 2: ", count)

}

func readIntoMatrix(fname string) [][]string {
	lines := utils.MustReadLines(fname)
	var result [][]string
	for _, line := range lines {
		var temp []string
		for _, char := range line {
			temp = append(temp, string(char))
		}
		result = append(result, temp)
	}
	return result
}

func stringGetter(matrix [][]string, c Point, d Point) string {
	Xmax := len(matrix[0])
	Ymax := len(matrix)
	remainingChar := ""
	for {
		if c.x < Xmax && c.y < Ymax && c.x >= 0 && c.y >= 0 {
			remainingChar += matrix[c.y][c.x]
			c.x += d.x
			c.y += d.y
		} else {
			break
		}
	}
	return remainingChar
}

func findNumberOfMatches(s string) int {
	p := `^XMAS`
	r, _ := regexp.Compile(p)
	matches := r.FindAllString(s, -1)
	if matches == nil {
		return 0
	} else {
		return 1
	}
}

func getCountsForLocation(matrix [][]string, p Point) int {
	ld := []Point{{1, 0}, {0, 1}, {1, 1}, {1, -1}, {-1, 0}, {0, -1}, {-1, -1}, {-1, 1}}
	count := 0
	for _, d := range ld {
		remainingChar := stringGetter(matrix, p, d)
		count += findNumberOfMatches(remainingChar)
	}
	return count
}

func isXMAS(matrix [][]string, p Point) bool {
	Xmax := len(matrix[0])
	Ymax := len(matrix)
	// precondtion to start should be M or S and have 2 more X direction and 2 more Y direction left
	if p.x > Xmax-3 || p.y > Ymax-3 || !(matrix[p.y][p.x] == "M" || matrix[p.y][p.x] == "S") {
		return false
	}

	string1 := matrix[p.y][p.x] + matrix[p.y+1][p.x+1] + matrix[p.y+2][p.x+2]
	string2 := matrix[p.y][p.x+2] + matrix[p.y+1][p.x+1] + matrix[p.y+2][p.x]
	if isMAS(string1) && isMAS(string2) {
		return true
	}
	return false

}

func isMAS(s string) bool {
	if s == "MAS" || s == "SAM" {
		return true
	}
	return false
}
