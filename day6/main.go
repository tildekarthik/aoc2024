package main

import (
	"aoc2024/utils"
	"fmt"
)

type coord struct {
	x int
	y int
}

type visitedSimilar struct {
	coord coord
	dir   coord
}

var (
	up    coord = coord{x: 0, y: -1}
	right coord = coord{x: 1, y: 0}
	down  coord = coord{x: 0, y: 1}
	left  coord = coord{x: -1, y: 0}
)

func main() {
	input := utils.ReadIntoMatrix("day6/input.txt")
	curr := getStartCoord(input)
	dir := up
	part1GetCovered(input, curr, dir)
	fmt.Println(part2IsLooping(input, curr, dir))
}

func part1GetCovered(input [][]string, curr coord, dir coord) {
	visitedPositions := make(map[coord]bool)
	maxX := len(input[0])
	maxY := len(input)
	for {
		visitedPositions[curr] = true
		nextStepCoord := moveOneStep(curr, dir)
		if isStepOutside(nextStepCoord, maxX, maxY) {
			break
		}
		nextStepChar := getValueAtCoord(nextStepCoord, input)
		dir = getNextStepDir(nextStepChar, dir)
		curr = moveOneStep(curr, dir)
	}
	println(len(visitedPositions))
}

func part2IsLooping(input [][]string, initpos coord, initdir coord) int {
	coords2PutHash := getCoords2PutHash(input)
	fmt.Println("Testing ", len(coords2PutHash), " coords")
	countLoopingPos := 0
	for _, c := range coords2PutHash {
		newInput := makeDeepCopy(input)
		newInput[c.y][c.x] = "#"
		if isLooping(newInput, initpos, initdir) {
			// fmt.Println("Looping at ", c)
			countLoopingPos++
		}
	}
	return countLoopingPos
}

func isLooping(input [][]string, curr coord, dir coord) bool {
	visitedPositions := make(map[coord]bool)
	visitedSimilarTimes := make(map[visitedSimilar]int)
	maxX := len(input[0])
	maxY := len(input)
	for {
		visitedPositions[curr] = true
		visitedSimilarTimes[visitedSimilar{coord: curr, dir: dir}]++
		if visitedSimilarTimes[visitedSimilar{coord: curr, dir: dir}] > 1 {
			return true
		}
		nextStepCoord := moveOneStep(curr, dir)
		if isStepOutside(nextStepCoord, maxX, maxY) {
			break
		}
		nextStepChar := getValueAtCoord(nextStepCoord, input)
		dir = getNextStepDir(nextStepChar, dir)
		if getValueAtCoord(moveOneStep(curr, dir), input) == "#" {
			dir = rotateRight(dir)
		}
		curr = moveOneStep(curr, dir)
	}
	return false
}

func rotateRight(c coord) coord {
	switch c {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	}
	return c
}

func getNextStepDir(nextChar string, currDir coord) coord {
	if nextChar == "#" {
		return rotateRight(currDir)
	} else {
		return currDir
	}
}

func moveOneStep(c coord, dir coord) coord {
	return coord{x: c.x + dir.x, y: c.y + dir.y}
}

func isStepOutside(c coord, maxX int, maxY int) bool {
	return c.x < 0 || c.x >= maxX || c.y < 0 || c.y >= maxY
}

func getStartCoord(input [][]string) coord {
	for y, row := range input {
		for x, char := range row {
			if char == "^" {
				return coord{x: x, y: y}
			}
		}
	}
	return coord{x: -1, y: -1}
}

func getValueAtCoord(c coord, input [][]string) string {
	return input[c.y][c.x]
}

func makeDeepCopy(input [][]string) [][]string {
	var result [][]string
	for _, row := range input {
		var temp []string
		temp = append(temp, row...)
		result = append(result, temp)
	}
	return result
}

func getCoords2PutHash(input [][]string) []coord {
	var result []coord
	maxX := len(input[0])
	maxY := len(input)
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if input[y][x] == "." {
				result = append(result, coord{x: x, y: y})
			}
		}
	}
	return result
}
