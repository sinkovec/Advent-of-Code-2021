#!/bin/bash

if [[ $# -eq 0 ]]; then
    latest=$(find . -maxdepth 1 -type d -name "day*" | sort | tail -1)
    latest=$(("${latest: -1}"+1))
else
    latest=$1
fi
latest=$(printf "%02d" $latest)
mkdir "day$latest"

# Generate main go file
echo "package day$latest

import (
	\"fmt\"
	\"strings\"

	\"github.com/sinkovec/Advent-of-Code-2021/util\"
)

func Run() {
	data := transform(util.ReadInput(\"day$latest/input.txt\"))
	fmt.Println(\"Day $latest\")
	fmt.Printf(\"\\tOne: %d\\n\", runOne(data))
	fmt.Printf(\"\\tTwo: %d\\n\", runTwo(data))
}

func runOne(data []string) int {
	return 0
}

func runTwo(data []string) int {
	return 0
}

func transform(data string) []string {
	result := make([]string, 0)
	for _, line := range strings.Split(data, \"\\n\") {
		result = append(result, line)
	}
	return result
}
" > ./day$latest/day$latest.go

# Generate test file
echo "package day$latest

import (
	\"testing\"

	\"github.com/stretchr/testify/assert\"
)

var data = []string {}

func Test_Day${latest}_One(t *testing.T) {
	assert.Equal(t, 0, runOne(data))
}

func Test_Day${latest}_Two(t *testing.T) {
	assert.Equal(t, 0, runTwo(data))
}

" > ./day$latest/day${latest}_test.go

# Create input.txt & README.md
touch ./day$latest/input.txt
touch ./day$latest/README.md

# Adjust main.go
sed -i "s/d \"github.com\/sinkovec\/Advent-of-Code-2021\/day.*\"/d \"github.com\/sinkovec\/Advent-of-Code-2021\/day$latest\"/" ./main.go