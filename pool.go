package pool

import "sync"

// Pool config
type Pool struct {
	concurrent int
	guard      chan struct{}
	wg         *sync.WaitGroup
}

// New create pool
func New(concurrent int) *Pool {
	return &Pool{
		concurrent: concurrent,
		guard:      make(chan struct{}, concurrent),
		wg:         &sync.WaitGroup{},
	}
}

// Go execute func
func (p *Pool) Go(fn func()) {
	p.guard <- struct{}{}
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer func() {
			<-p.guard
		}()
		fn()
	}()
}

// Wait wait execute routine
func (p *Pool) Wait() {
	p.wg.Wait()
}
