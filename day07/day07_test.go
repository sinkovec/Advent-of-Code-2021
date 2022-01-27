package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []int{
	16, 1, 2, 0, 4, 2, 7, 1, 2, 14,
}

func Test_Day07_One(t *testing.T) {
	assert.Equal(t, 37, runOne(data))
}

func Test_Day07_Two(t *testing.T) {
	assert.Equal(t, 168, runTwo(data))
}
