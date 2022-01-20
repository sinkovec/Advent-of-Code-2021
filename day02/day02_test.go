package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data = []command{
	{"forward", 5},
	{"down", 5},
	{"forward", 8},
	{"up", 3},
	{"down", 8},
	{"forward", 2},
}

func Test_Day02_One(t *testing.T) {
	assert.Equal(t, 150, runOne(data))
}

func Test_Day02_Two(t *testing.T) {
	assert.Equal(t, 900, runTwo(data))
}
