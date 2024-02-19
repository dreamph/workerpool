package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/dreamph/workerpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	pool := workerpool.NewPool(ctx, 20, 10)

	// Submit 100 tasks
	for i := 0; i < 100; i++ {
		n := i
		pool.Submit(func() {
			TestFn(n)
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	pool.Wait()

}

func TestFn(n int) {
	time.Sleep(3 * time.Second)
}
