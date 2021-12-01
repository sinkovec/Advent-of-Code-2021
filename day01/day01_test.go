package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func Test_Day01_One(t *testing.T) {
	assert.Equal(t, 7, runOne(data))
}

func Test_Day01_Two(t *testing.T) {
	assert.Equal(t, 5, runTwo(data))
}
