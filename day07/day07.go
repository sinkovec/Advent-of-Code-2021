package day07

import (
	"fmt"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

func Run() {
	data := util.ReadInts("day07/input.txt")
	fmt.Println("Day 07")
	fmt.Printf("\tOne: %d\n", runOne(data))
	fmt.Printf("\tTwo: %d\n", runTwo(data))
}

func runOne(data []int) int {
	maxPos := util.Max(data)
	result := -1
	var fuel int
	for targetPos := 1; targetPos < maxPos; targetPos++ {
		fuel = 0
		for _, crabPos := range data {
			fuel += util.Abs(crabPos - targetPos)
		}
		if result < 0 || fuel < result {
			result = fuel
		}
	}
	return result
}

func runTwo(data []int) int {
	maxPos := util.Max(data)
	result := -1
	var fuel int
	var distance int
	for targetPos := 1; targetPos < maxPos; targetPos++ {
		fuel = 0
		for _, crabPos := range data {
			distance = util.Abs(crabPos - targetPos)
			fuel += (distance * (distance + 1)) / 2
		}
		if result < 0 || fuel < result {
			result = fuel
		}
	}
	return result
}
