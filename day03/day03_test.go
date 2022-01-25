package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = [][]int{
	{0, 0, 1, 0, 0},
	{1, 1, 1, 1, 0},
	{1, 0, 1, 1, 0},
	{1, 0, 1, 1, 1},
	{1, 0, 1, 0, 1},
	{0, 1, 1, 1, 1},
	{0, 0, 1, 1, 1},
	{1, 1, 1, 0, 0},
	{1, 0, 0, 0, 0},
	{1, 1, 0, 0, 1},
	{0, 0, 0, 1, 0},
	{0, 1, 0, 1, 0},
}

func Test_Day03_One(t *testing.T) {
	assert.Equal(t, 198, runOne(data))
}

func Test_Day03_Two(t *testing.T) {
	assert.Equal(t, 230, runTwo(data))
}
