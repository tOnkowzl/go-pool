package pool

import (
	"sync"
)

type fn func()

// Pool config
type Pool struct {
	guard chan struct{}
	wg    sync.WaitGroup
}

// New create pool
func New(concurrent int) *Pool {
	return &Pool{
		guard: make(chan struct{}, concurrent),
		wg:    sync.WaitGroup{},
	}
}

// Go execute func
func (p *Pool) Go(fn fn) {
	p.guard <- struct{}{}
	p.wg.Add(1)
	go p.execute(fn)
}

func (p *Pool) execute(fn fn) {
	fn()
	p.wg.Done()
	<-p.guard
}

// Wait execute routine
func (p *Pool) Wait() {
	p.wg.Wait()
}
