package pool_test

import (
	"fmt"
	"testing"

	"github.com/tOnkowzl/go-pool"
)

func TestGo(t *testing.T) {
	p := pool.New(10)

	for i := 1; i <= 1000; i++ {
		arg := i
		p.Go(func() error {
			fmt.Println(arg)
			return nil
		})
	}

	p.Wait()
}

func Benchmark_Go(b *testing.B) {
	p := pool.New(1)

	for i := 1; i <= b.N; i++ {
		p.Go(func() error {
			return nil
		})
	}

	p.Wait()
}
