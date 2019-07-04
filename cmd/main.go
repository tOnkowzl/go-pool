package main

import (
	"fmt"

	"github.com/tOnkowzl/go-pool"
)

func main() {
	p := pool.New(10)

	for i := 1; i <= 1000; i++ {
		p.Go(func(args ...interface{}) {
			fmt.Println(args[0])
		}, i)
	}

	p.Wait()
}
