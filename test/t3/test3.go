package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(1)

	go func() {
		c1 := make(chan os.Signal, 1)
		signal.Notify(c1, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		fmt.Printf("goroutine 1 receive a signal : %v\n\n", <-c1)
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("all groutine done!\n")
}
