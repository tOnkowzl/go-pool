# go-pool

## Benchmark

```text
Benchmark_Go-8   	16753122	        74.2 ns/op	       0 B/op	       0 allocs/op
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
		p.Go(func(){
			fmt.Println(arg)
		})
	}

	p.Wait()
}
```
