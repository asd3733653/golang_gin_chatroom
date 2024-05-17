package main

import (
	"fmt"
)

func main() {
	q := make(chan int, 4)
	q <- 1
	select {
	case val := <-q:
		fmt.Printf("ok: %d\n", val)
	default:
		fmt.Println("wrong")
	}

	q <- 2

	ass := <-q
	fmt.Printf("%d\n", ass)
}
