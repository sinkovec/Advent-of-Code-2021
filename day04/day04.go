package day04

import (
	"fmt"
	"strings"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

func Run() {
	data := transform(util.ReadInput("day04/input.txt"))
	fmt.Println("Day 04")
	fmt.Printf("\tOne: %d\n", runOne(data))
	fmt.Printf("\tTwo: %d\n", runTwo(data))
}

func runOne(data []string) int {
	return 0
}

func runTwo(data []string) int {
	return 0
}

func transform(data string) []string {
	result := make([]string, 0)
	for _, line := range strings.Split(data, "\n") {
		result = append(result, line)
	}
	return result
}
