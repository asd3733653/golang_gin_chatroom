package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var a int = 1
	fmt.Print(a)
	ch := make(chan struct{}, 5)
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int) {
			fmt.Printf("i : %d\n", i)
			defer wg.Done()
			<-ch
		}(i)
	}
	wg.Wait()
}
