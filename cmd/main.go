package main

import (
	"fmt"

	"github.com/tOnkowzl/go-pool"
)

func main() {
	p := pool.New(10)

	for i := 1; i <= 1000; i++ {
		p.Go(func(i int) func() {
			return func() {
				fmt.Println(i)
			}
		}(i))
	}

	p.Wait()
}
