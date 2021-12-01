package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

func Run() {
	data := transform(util.ReadInput("day01/input.txt"))
	fmt.Println("Day 01")
	fmt.Printf("\tOne: %d\n", runOne(data))
	fmt.Printf("\tTwo: %d\n", runTwo(data))
}

func runOne(data []int) int {
	numInc := 0
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			numInc++
		}
	}
	return numInc
}

func runTwo(data []int) int {
	numInc := 0
	for i := 3; i < len(data); i++ {
		val1 := data[i-3] + data[i-2] + data[i-1]
		val2 := data[i-2] + data[i-1] + data[i]
		if val1 < val2 {
			numInc++
		}
	}
	return numInc
}

func transform(data string) []int {
	result := make([]int, 0)
	for _, s := range strings.Split(data, "\n") {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}
