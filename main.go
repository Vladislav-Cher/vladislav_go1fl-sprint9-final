package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 || size > 100_000_000 {
		fmt.Printf("incorrect amount of numbers: %d\n", size)
		return nil
	}
	res := make([]int, size)
	for i := 0; i < size; i++ {
		num := rand.Intn(size)
		res[i] = num
	}
	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		fmt.Print("data is empty\n")
		return 0
	}
	number := data[0]
	for _, n := range data {
		if n > number {
			number = n
		}
	}
	return number
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		fmt.Print("data is empty\n")
		return 0
	}
	var wg sync.WaitGroup
	maxOfMax := make([]int, CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			beginningOfChunk := index * (len(data) / CHUNKS)
			// из любопытства пробовал менять CHUNKS, поэтому добавил следующие строки
			// на случай если CHUNKS не будет делить SIZE без остатка.
			// *эту задачу на моем ПК выполняют быстрее всего 5 горутин
			var endOfChunk int
			if index == CHUNKS-1 {
				endOfChunk = len(data)
			} else {
				endOfChunk = index*(len(data)/CHUNKS) + (len(data) / CHUNKS)
			}

			number := maximum(data[beginningOfChunk:endOfChunk])
			maxOfMax[index] = number
		}(i)

	}
	wg.Wait()
	return maximum(maxOfMax)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	end := time.Now()
	elapsed := time.Duration(end.Sub(start).Milliseconds())

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	end = time.Now()
	elapsed = time.Duration(end.Sub(start).Milliseconds())

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
