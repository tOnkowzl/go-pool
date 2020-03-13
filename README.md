# go-pool

## Benchmark

```text
Benchmark_Go-8   	 3287200	       373 ns/op	       0 B/op	       0 allocs/op
```

## Example

```golang
package main

import (
	"fmt"

	"github.com/tOnkowzl/go-pool"
)

func main() {
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
```
