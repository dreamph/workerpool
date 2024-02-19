package workerpool

import (
	"github.com/alitto/pond"
	"golang.org/x/net/context"
)

type Pool struct {
	pool *pond.WorkerPool
}

func NewPool(ctx context.Context, maxWorkers, maxCapacity int) *Pool {
	return &Pool{pool: pond.New(maxWorkers, maxCapacity, pond.Context(ctx))}
}

func (p *Pool) Submit(task func()) {
	p.pool.Submit(task)
}

func (p *Pool) Wait() {
	p.pool.StopAndWait()
}
