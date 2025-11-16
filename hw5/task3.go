package main

import(
	"fmt"
	"sync"
)

func main(){
	var sm sync.Map
	var wg sync.WaitGroup

	for i:=0; i < 10; i++{
		wg.Add(1)
		go func (id int)  {
			defer wg.Done()
			key := fmt.Sprintf("key_%d", id)
			value := id*10
			sm.Store(key, value)
			fmt.Printf("Горутина %d записала: %s -> %d\n", id, key, value)
		}(i)
	}

	wg.Wait()

	fmt.Println("\nРезультаты:")
	for i := 0; i < 10; i++{
		key := fmt.Sprintf("key_%d", i)
		if value, ok := sm.Load(key); ok{
			fmt.Printf("%s: %d\n", key, value)
		}
	}
}