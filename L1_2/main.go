package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup

	// Запускаем горутину для каждого числа
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(num)
	}

	// Запускаем горутину для закрытия канала после завершения вычислений
	go func() {
		wg.Wait()
		close(results)
	}()

	for square := range results {
		fmt.Println(square)
	}
}
