package main

import (
	"fmt"
	"time"
)

func main() {
	c := []string{"apple", "banana", "cherry"}

	for _, v := range c {
		// go func(s string) {
		fmt.Println(v)
		// }(v)
	}

	time.Sleep(time.Second) // 等待協程執行完成
}
