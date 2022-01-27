package day06

import (
	"fmt"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

func Run() {
	data := util.ReadInts("day06/input.txt")
	fmt.Println("Day 06")
	fmt.Printf("\tOne: %d\n", runOne(data))
	fmt.Printf("\tTwo: %d\n", runTwo(data))
}

func runOne(data []int) int {
	return run(data, 80)
}

func runTwo(data []int) int {
	return run(data, 256)
}

func run(data []int, days int) int {
	const TIMER_MAX = 9
	timer := make([]int, 9)
	for _, v := range data {
		timer[v]++
	}
	for i := 0; i < days; i++ {
		newFishScore := timer[0]
		for j := 1; j < TIMER_MAX; j++ {
			timer[j-1] = timer[j]
		}
		timer[6] += newFishScore
		timer[8] = newFishScore
	}
	return util.Sum(timer)
}
