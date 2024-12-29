## Basic Usage


# Example Worker Pool
```go
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

type Result struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	pool := workerpool.NewResultPool[int, Result](ctx, 20, 10)

	// Submit 20 tasks
	for i := 0; i < 20; i++ {
		n := i
		pool.Submit(n, func() (Result, error) {
			return TestFn(n)
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	poolResponse := pool.Wait()

	poolResults := poolResponse.Result()

	for key, value := range poolResults {
		fmt.Println(fmt.Sprintf("%d : [%v,%s]", key, value.Data.Success, value.Data.Data))
	}

	// Output:
	/*
		0 : [true,Result of task #0]
		2 : [true,Result of task #2]
		15 : [true,Result of task #15]
		3 : [true,Result of task #3]
		12 : [true,Result of task #12]
		8 : [true,Result of task #8]
		6 : [true,Result of task #6]
		11 : [true,Result of task #11]
		1 : [true,Result of task #1]
		9 : [true,Result of task #9]
		18 : [true,Result of task #18]
		4 : [true,Result of task #4]
		13 : [true,Result of task #13]
		19 : [true,Result of task #19]
		10 : [true,Result of task #10]
		7 : [true,Result of task #7]
		14 : [true,Result of task #14]
		16 : [true,Result of task #16]
		17 : [true,Result of task #17]
		5 : [true,Result of task #5]
	*/
}

func TestFn(n int) (Result, error) {
	v := fmt.Sprintf("Result of task #%d", n)
	//fmt.Println(v)
	time.Sleep(3 * time.Second)
	return Result{Success: true, Data: v}, nil
}


```