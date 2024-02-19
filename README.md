## Basic Usage


# Example Worker Pool
```go
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


```
# Example Worker Pool with Result
```go
package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"fmt"

	"github.com/dreamph/workerpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	pool := workerpool.NewResultPool[int, string](ctx, 20, 10)

	// Submit 100 tasks
	for i := 0; i < 100; i++ {
		n := i
		pool.Submit(n, func() (string, error) {
			return TestFn(n)
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	poolResponse := pool.Wait()

	poolResult := poolResponse.Result()

	// Get result from first job
	fmt.Println(poolResult[0])

	fmt.Println(len(poolResult))
}

func TestFn(n int) (string, error) {
	v := fmt.Sprintf("Running task #%d\n", n)
	//fmt.Println(v)
	time.Sleep(3 * time.Second)
	return v, nil
}

```