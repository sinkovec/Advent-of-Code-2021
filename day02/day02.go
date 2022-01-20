package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

type command struct {
	direction string
	units     int
}

func Run() {
	data := transform(util.ReadInput("day02/input.txt"))
	fmt.Println("Day 02")
	fmt.Printf("\tOne: %d\n", runOne(data))
	fmt.Printf("\tTwo: %d\n", runTwo(data))
}

func runOne(data []command) int {
	hpos := 0
	depth := 0
	for _, cmd := range data {
		if cmd.direction == "forward" {
			hpos += cmd.units
		} else if cmd.direction == "down" {
			depth += cmd.units
		} else if cmd.direction == "up" {
			depth -= cmd.units
		}
	}
	return hpos * depth
}

func runTwo(data []command) int {
	hpos := 0
	depth := 0
	aim := 0
	for _, cmd := range data {
		if cmd.direction == "forward" {
			hpos += cmd.units
			depth += aim * cmd.units
		} else if cmd.direction == "down" {
			aim += cmd.units
		} else if cmd.direction == "up" {
			aim -= cmd.units
		}
	}
	return hpos * depth
}

func transform(data string) []command {
	result := make([]command, 0)
	for _, line := range strings.Split(data, "\n") {
		s := strings.Split(line, " ")
		i, _ := strconv.Atoi(s[1])
		item := command{s[0], i}
		result = append(result, item)
	}
	return result
}
