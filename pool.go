package pool

import (
	"golang.org/x/sync/errgroup"
)

type fn func() error

// Pool config
type Pool struct {
	guard chan struct{}
	eg    errgroup.Group
}

// New create pool
func New(concurrent int) *Pool {
	return &Pool{
		guard: make(chan struct{}, concurrent),
		eg:    errgroup.Group{},
	}
}

// Go execute func
func (p *Pool) Go(fn fn) {
	p.guard <- struct{}{}
	p.eg.Go(p.execute(fn))
}

func (p *Pool) execute(fn fn) fn {
	defer func() {
		<-p.guard
	}()
	return fn
}

// Wait execute routine
func (p *Pool) Wait() error {
	return p.eg.Wait()
}
