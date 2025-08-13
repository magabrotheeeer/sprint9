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

func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]int, size)
	for i := range size {
		res[i] = rnd.Int()
	}
	return res
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	maxNumber := data[0]
	for _, v := range data {
		if v > maxNumber {
			maxNumber = v
		}
	}
	return maxNumber
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	sizePart := len(data) / CHUNKS
	maxResults := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := range CHUNKS {
		start := i * sizePart
		end := start + sizePart
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(idx int, part []int) {
			defer wg.Done()
			maxResults[idx] = maximum(part)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxResults)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	randomElements := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	startOneRoutine := time.Now().UTC()
	max := maximum(randomElements)
	endOneRoutine := time.Now().UTC().Sub(startOneRoutine).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, endOneRoutine)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	startRoutines := time.Now().UTC()
	max = maxChunks(randomElements)
	endRoutines := time.Now().UTC().Sub(startRoutines).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, endRoutines)
}
