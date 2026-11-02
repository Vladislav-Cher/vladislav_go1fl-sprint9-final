package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	sizes := []int{-1, 0, 1, 10, 100_000_000, 100_100_100}
	for i, size := range sizes {
		res := generateRandomElements(size)
		if size > 0 && size <= 100_000_000 {
			require.Equal(t, sizes[i], len(res))
		} else {
			require.Nil(t, res)
		}
	}
}

func TestMaximum(t *testing.T) {
	//если слайс чисел == nil
	var data []int
	res := maximum(data)
	require.Equal(t, 0, res)

	// если слайс чисел == 0
	data = []int{}
	res = maximum(data)
	require.Equal(t, 0, res)

	// если в слайсе есть отрицательные и положительные числа
	data = []int{-528, -5, 0, 10, 428}
	res = maximum(data)
	require.Equal(t, 428, res)

	// если в слайсе только отрицательные числа
	data = []int{-528, -5, -10, -428}
	res = maximum(data)
	require.Equal(t, -5, res)

	// если в слайсе только одно число
	data = []int{5}
	res = maximum(data)
	require.Equal(t, 5, res)
}
