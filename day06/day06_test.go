package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []int{
	3, 4, 3, 1, 2,
}

func Test_Day06_One(t *testing.T) {
	assert.Equal(t, 5934, runOne(data))
}

func Test_Day06_Two(t *testing.T) {
	assert.Equal(t, 26984457539, runTwo(data))
}
