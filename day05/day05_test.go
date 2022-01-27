package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []Line{
	{Point{0, 9}, Point{5, 9}},
	{Point{8, 0}, Point{0, 8}},
	{Point{9, 4}, Point{3, 4}},
	{Point{2, 2}, Point{2, 1}},
	{Point{7, 0}, Point{7, 4}},
	{Point{6, 4}, Point{2, 0}},
	{Point{0, 9}, Point{2, 9}},
	{Point{3, 4}, Point{1, 4}},
	{Point{0, 0}, Point{8, 8}},
	{Point{5, 5}, Point{8, 2}},
}

func Test_Day05_One(t *testing.T) {
	assert.Equal(t, 5, runOne(data))
}

func Test_Day05_Two(t *testing.T) {
	assert.Equal(t, 12, runTwo(data))
}
