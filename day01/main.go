package day01

import (
	"bufio"
	"log"
	"os"
)

func run_one() int {
	return 1
}

func run_two() int {
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
