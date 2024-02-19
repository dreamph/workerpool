package workerpool

import (
	"github.com/alitto/pond"
	"golang.org/x/net/context"
)

type Result[T any] struct {
	Data T     `json:"key"`
	Err  error `json:"value"`
}

type ResultPool[K comparable, V any] struct {
	pool      *pond.WorkerPool
	resultMap Map[K, Result[V]]
}

func NewResultPool[K comparable, V any](ctx context.Context, maxWorkers, maxCapacity int) *ResultPool[K, V] {
	return &ResultPool[K, V]{pool: pond.New(maxWorkers, maxCapacity, pond.Context(ctx))}
}

func (p *ResultPool[K, V]) Submit(key K, task func() (V, error)) {
	p.pool.Submit(func() {
		p.doTask(key, task)
	})
}

func (p *ResultPool[K, V]) Wait() *Response[K, V] {
	p.pool.StopAndWait()

	result := map[K]Result[V]{}
	p.resultMap.Range(func(key K, value Result[V]) bool {
		result[key] = value
		return true
	})
	return NewResponse[K, V](&p.resultMap)
}

func (p *ResultPool[K, V]) doTask(key K, task func() (V, error)) {
	taskResult, err := task()
	if err != nil {
		p.resultMap.Store(key, Result[V]{
			Err: err,
		})
	} else {
		p.resultMap.Store(key, Result[V]{
			Data: taskResult,
		})
	}
}

type Response[K comparable, V any] struct {
	dataMap *Map[K, Result[V]]
}

func NewResponse[K comparable, V any](resultMap *Map[K, Result[V]]) *Response[K, V] {
	return &Response[K, V]{
		dataMap: resultMap,
	}
}

func (r *Response[K, V]) Result() map[K]Result[V] {
	result := map[K]Result[V]{}
	r.dataMap.Range(func(key K, value Result[V]) bool {
		result[key] = value
		return true
	})
	return result
}
