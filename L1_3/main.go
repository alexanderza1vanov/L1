package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case v, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("[worker %d] got: %d\n", id, v)
		}
	}
}

func main() {
	n := flag.Int("n", 4, "number of workers")
	flag.Parse()

	jobs := make(chan int, 128)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		cancel()
		close(jobs) // прекращаем выдачу задач
	}()

	// запускаем N воркеров
	var wg sync.WaitGroup
	wg.Add(*n)
	for i := 1; i <= *n; i++ {
		go worker(ctx, i, jobs, &wg)
	}

	// главный продюсер: постоянно пишет в канал
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	counter := 0
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			return
		case <-ticker.C:
			counter++
			select {
			case jobs <- counter:
			case <-ctx.Done():
				wg.Wait()
				return
			}
		}
	}
}
