package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	nums:= []int{2,4,6,8,10}
	var sum int64
	var wg sync.WaitGroup

	for _, num := range nums{
		wg.Add(1)
		go func (n int)  {
			defer wg.Done()
			atomic.AddInt64(&sum, int64(n*n))
		}(num)
	}
	wg.Wait()
	fmt.Printf("Сумма квадратов: %d\n", sum)
}