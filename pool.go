package pool

import "sync"

type fn func(args []interface{})

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
func (p *Pool) Go(fn fn, args ...interface{}) {
	p.guard <- struct{}{}
	p.wg.Add(1)
	go p.execute(fn, args)
}

func (p *Pool) execute(fn fn, args []interface{}) {
	defer p.wg.Done()
	defer func() {
		<-p.guard
	}()
	fn(args)
}

// Wait execute routine
func (p *Pool) Wait() {
	p.wg.Wait()
}
