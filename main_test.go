package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var options = []struct {
	name string
	data []int
	want int
}{
	{"nil slice", nil, 0},
	{"empty slice", []int{}, 0},
	{"mixed numbers", []int{-528, -5, 0, 10, 428}, 428},
	{"all negatives", []int{-528, -5, -10, -428}, -5},
	{"single element", []int{5}, 5},
}

func TestGenerateRandomElements(t *testing.T) {
	sizes := []int{-1, 0, 1, 10, 100_000_000, 100_100_100}
	for i, size := range sizes {
		res := generateRandomElements(size)
		if size > 0 {
			assert.Equal(t, sizes[i], len(res))
		} else {
			assert.Nil(t, res)
		}
	}
}

func TestMaximum(t *testing.T) {
	for _, option := range options {
		t.Run(option.name, func(t *testing.T) {
			assert.Equal(t, option.want, maximum(option.data), "for test %s", option.name)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	for _, option := range options {
		t.Run(option.name, func(t *testing.T) {
			assert.Equal(t, option.want, maxChunks(option.data), "for test %s", option.name)
		})
	}
}
