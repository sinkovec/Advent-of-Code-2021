package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Day01_One(t *testing.T) {
	assert.Equal(t, 1, run_one())
}

func Test_Day01_Two(t *testing.T) {
	assert.Equal(t, 2, run_two())
}
