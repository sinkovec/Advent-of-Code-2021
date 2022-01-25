package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sinkovec/Advent-of-Code-2021/util"
)

func Run() {
	nums, boards := transform(util.ReadInput("day04/input.txt"))
	fmt.Println("Day 04")
	fmt.Printf("\tOne: %d\n", runOne(nums, boards))
	fmt.Printf("\tTwo: %d\n", runTwo(nums, boards))
}

func runOne(nums []int, boards [][][]int) int {
	N := len(boards[0])

	m, c, marked := transformToMap(boards)

	winningBoard := -1
	lastNum := -1
	for _, num := range nums {
		for _, info := range m[num] {
			k, i, j := info[0], info[1], info[2]
			c[k][i]++
			c[k][j+N]++
			marked[k][i][j] = true
			if c[k][i] == 5 || c[k][j+N] == 5 {
				winningBoard = k
				lastNum = num
			}
		}
		if winningBoard != -1 {
			break
		}
	}
	return result(boards, marked, winningBoard, lastNum)
}

func runTwo(nums []int, boards [][][]int) int {
	K := len(boards)
	N := len(boards[0])

	m, c, marked := transformToMap(boards)

	winningBoards := make([]bool, 0)
	for k := 0; k < K; k++ {
		winningBoards = append(winningBoards, false)
	}
	losingBoard := -1
	lastNum := -1
	for _, num := range nums {
		for _, info := range m[num] {
			k, i, j := info[0], info[1], info[2]
			c[k][i]++
			c[k][j+N]++
			marked[k][i][j] = true
			if c[k][i] == 5 || c[k][j+N] == 5 {
				winningBoards[k] = true
				nTrue := 0
				for _, v := range winningBoards {
					if v {
						nTrue++
					}
				}
				if nTrue == K {
					losingBoard = k
					lastNum = num
					break
				}
			}
		}
		if losingBoard != -1 {
			break
		}
	}
	return result(boards, marked, losingBoard, lastNum)
}

func result(boards [][][]int, marked [][][]bool, k int, lastNum int) int {
	N := len(boards[0])
	M := len(boards[0][0])
	result := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if !marked[k][i][j] {
				result += boards[k][i][j]
			}
		}
	}
	return result * lastNum
}

func transformToMap(boards [][][]int) (map[int][][]int, [][]int, [][][]bool) {
	// transform board to map structure: for each value, store a list of triples (k, i, j) (value is stored in cell (i,j) of board k)
	K := len(boards)
	N := len(boards[0])
	M := len(boards[0][0])
	m := make(map[int][][]int)
	c := make([][]int, 0)
	marked := make([][][]bool, 0)
	for k := 0; k < K; k++ {
		marked = append(marked, make([][]bool, 0))
		for i := 0; i < N; i++ {
			marked[k] = append(marked[k], make([]bool, 0))
			for j := 0; j < M; j++ {
				m[boards[k][i][j]] = append(m[boards[k][i][j]], []int{k, i, j})
				marked[k][i] = append(marked[k][i], false)
			}
		}
	}
	// initialize counting map
	for k := 0; k < K; k++ {
		c = append(c, make([]int, 0))
		for i := 0; i < N+M; i++ {
			c[k] = append(c[k], 0)
		}
	}
	return m, c, marked
}

func transform(data string) ([]int, [][][]int) {
	nums := make([]int, 0)
	boards := make([][][]int, 0)
	first := true
	k := -1
	i := 0
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			boards = append(boards, make([][]int, 0))
			i = 0
			k++
			continue
		}
		if first {
			first = false
			for _, x := range strings.Split(line, ",") {
				num, _ := strconv.Atoi(x)
				nums = append(nums, num)
			}
		} else {
			boards[k] = append(boards[k], make([]int, 0))
			for _, x := range strings.Split(line, " ") {
				if x == "" {
					continue
				}
				num, _ := strconv.Atoi(x)
				boards[k][i] = append(boards[k][i], num)
			}
			i++
		}
	}
	return nums, boards
}
