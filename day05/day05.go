package day05

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func Run() {
	data := transform(util.ReadInput("day05/input.txt"))
	fmt.Println("Day 05")
	fmt.Printf("\tOne: %d\n", runOne(data))
	fmt.Printf("\tTwo: %d\n", runTwo(data))
}

func runOne(data []Line) int {
	vents := getVentMap(data)
	for _, line := range data {
		start, end := line.start, line.end
		if start.x != end.x && start.y != end.y {
			continue
		}
		for i := start.y; i <= end.y; i++ {
			for j := start.x; j <= end.x; j++ {
				vents[i][j]++
			}
		}
	}
	return result(vents)
}

func runTwo(data []Line) int {
	vents := getVentMap(data)
	for _, line := range data {
		start, end := line.start, line.end
		if start.x != end.x && start.y != end.y {
			if start.y < end.y {
				for i := 0; i <= end.x-start.x; i++ {
					vents[start.y+i][start.x+i]++
				}
			} else {
				for i := 0; i <= end.x-start.x; i++ {
					vents[start.y-i][start.x+i]++
				}
			}
		} else {
			for i := start.y; i <= end.y; i++ {
				for j := start.x; j <= end.x; j++ {
					vents[i][j]++
				}
			}
		}
	}
	return result(vents)
}

func getVentMap(data []Line) [][]int {
	data = sortLines(data)
	xMax, yMax := getMax(data)
	xMax++
	yMax++
	result := make([][]int, yMax)
	for i := range result {
		result[i] = make([]int, xMax)
	}
	return result
}

func sortLines(data []Line) []Line {
	var tmp int
	for i := 0; i < len(data); i++ {
		if data[i].start.x == data[i].end.x {
			if data[i].start.y > data[i].end.y {
				tmp = data[i].start.y
				data[i].start.y = data[i].end.y
				data[i].end.y = tmp
			}
		} else if data[i].start.y == data[i].end.y {
			if data[i].start.x > data[i].end.x {
				tmp = data[i].start.x
				data[i].start.x = data[i].end.x
				data[i].end.x = tmp
			}
		} else if data[i].start.x > data[i].end.x {
			tmp = data[i].start.x
			data[i].start.x = data[i].end.x
			data[i].end.x = tmp
			tmp = data[i].start.y
			data[i].start.y = data[i].end.y
			data[i].end.y = tmp
		}
	}
	return data
}

func getMax(data []Line) (int, int) {
	xMax, yMax := -1, -1
	for _, line := range data {
		if line.end.x > xMax {
			xMax = line.end.x
		}
		if line.end.y > yMax {
			yMax = line.end.y
		}
	}
	return xMax, yMax
}

func result(vents [][]int) int {
	result := 0
	for i := 0; i < len(vents); i++ {
		for j := 0; j < len(vents[i]); j++ {
			if vents[i][j] >= 2 {
				result++
			}
		}
	}
	return result
}

func transform(data string) []Line {
	result := make([]Line, 0)
	for _, line := range strings.Split(data, "\n") {
		coordinates := strings.Split(line, " -> ")
		start, end := coordinates[0], coordinates[1]
		result = append(result, Line{toPoint(start), toPoint(end)})
	}
	return result
}

func toPoint(s string) Point {
	coordinate := strings.Split(s, ",")
	x, _ := strconv.Atoi(coordinate[0])
	y, _ := strconv.Atoi(coordinate[1])
	return Point{x, y}
}
