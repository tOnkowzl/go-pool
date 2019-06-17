package pool

import "sync"

// Pool config
type Pool struct {
	guard chan struct{}
	wg    *sync.WaitGroup
}

// New create pool
func New(concurrent int) *Pool {
	return &Pool{
		guard: make(chan struct{}, concurrent),
		wg:    &sync.WaitGroup{},
	}
}

// Go execute func
func (p *Pool) Go(fn func()) {
	p.guard <- struct{}{}
	p.wg.Add(1)
	go p.execute(fn)
}

func (p *Pool) execute(fn func()) {
	defer p.wg.Done()
	defer func() {
		<-p.guard
	}()
	fn()
}

// Wait wait execute routine
func (p *Pool) Wait() {
	p.wg.Wait()
}
