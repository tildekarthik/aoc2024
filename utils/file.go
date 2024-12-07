package utils

import (
	"bufio"
	"os"
	"strconv"
)

func MustReadLines(fname string) []string {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var lines []string
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines
}

func MustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func ReadIntoMatrix(fname string) [][]string {
	lines := MustReadLines(fname)
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
