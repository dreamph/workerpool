package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/dreamph/workerpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	pool := workerpool.NewPool(ctx, 20, 10)

	// Submit 20 tasks
	for i := 0; i < 20; i++ {
		n := i
		pool.Submit(func() {
			TestFn(n)
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	pool.Wait()

	// Output:
	/*
		End of task #0
		End of task #10
		End of task #5
		End of task #13
		End of task #8
		End of task #4
		End of task #9
		End of task #11
		End of task #12
		End of task #19
		End of task #7
		End of task #15
		End of task #16
		End of task #17
		End of task #6
		End of task #18
		End of task #3
		End of task #1
		End of task #14
		End of task #2
	*/

}

func TestFn(n int) {
	time.Sleep(3 * time.Second)
	fmt.Println(fmt.Sprintf("End of task #%d", n))
}
