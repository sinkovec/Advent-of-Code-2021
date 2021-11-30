package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Run() {
	fmt.Println("Day 01")
	fmt.Printf("\tOne: %d\n", runOne())
	fmt.Printf("\tTwo: %d\n", runTwo())
}

func runOne() int {
	return 1
}

func runTwo() int {
	return 2
}

func readInput(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var res string

	s := bufio.NewScanner(f)
	for s.Scan() {
		// do stuff
	}
	return res, s.Err()
}
