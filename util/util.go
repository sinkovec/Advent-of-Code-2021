package util

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadInput(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("could not open %s: %v", path, err)
	}
	return string(data)
}

func ReadIntPerLine(path string) []int {
	return ReadIntWithDelimiter(path, "\n")
}

func ReadInts(path string) []int {
	return ReadIntWithDelimiter(path, ",")
}

func ReadIntWithDelimiter(path string, delimiter string) []int {
	data := ReadInput(path)
	result := make([]int, 0)
	for _, s := range strings.Split(data, delimiter) {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}
