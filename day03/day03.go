package day03

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

func Run() {
	data := transform(util.ReadInput("day03/input.txt"))
	fmt.Println("Day 03")
	fmt.Printf("\tOne: %d\n", runOne(data))
	fmt.Printf("\tTwo: %d\n", runTwo(data))
}

func runOne(data [][]int) int {
	gamma := ""
	epsilon := ""
	for j := 0; j < len(data[0]); j++ {
		if mostCommon(data, j) == 1 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaInt, _ := strconv.ParseInt(gamma, 2, 0)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 0)
	return int(gammaInt) * int(epsilonInt)
}

func runTwo(data [][]int) int {
	oxygenGeneratorRating := data
	co2ScrubberRating := data
	return filter(oxygenGeneratorRating, mostCommon) * filter(co2ScrubberRating, leastCommon)
}

func filter(data [][]int, f func([][]int, int) int) int {
	yl := len(data[0])
	for j := 0; j < yl; j++ {
		m := f(data, j)
		filtered := make([][]int, 0)
		for i := 0; i < len(data); i++ {
			if data[i][j] == m {
				filtered = append(filtered, data[i])
			}
		}
		data = filtered
		if len(data) == 1 {
			break
		}
	}
	asString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(data[0])), ""), "[]")
	result, _ := strconv.ParseInt(asString, 2, 0)
	return int(result)
}

func mostCommon(data [][]int, col int) int {
	ones := 0
	zeros := 0
	for i := 0; i < len(data); i++ {
		if data[i][col] == 1 {
			ones++
		} else {
			zeros++
		}
	}
	if ones >= zeros {
		return 1
	} else {
		return 0
	}
}

func leastCommon(data [][]int, col int) int {
	if mostCommon(data, col) == 1 {
		return 0
	} else {
		return 1
	}
}

func transform(data string) [][]int {
	result := make([][]int, 0)
	for _, line := range strings.Split(data, "\n") {
		bits := make([]int, 0)
		for _, s := range strings.Split(line, "") {
			b, _ := strconv.Atoi(s)
			bits = append(bits, b)
		}
		result = append(result, bits)
	}
	return result
}
