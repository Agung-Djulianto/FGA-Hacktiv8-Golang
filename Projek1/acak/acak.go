package main

import (
	"fmt"
	"sync"
)

func main() {
	data1 := []string{"coba1", "coba2", "coba3"}
	data2 := []string{"bisa1", "bisa2", "bisa3"}

	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			fmt.Println(data1, i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Println(data2, i)
		}(i)
	}

	wg.Wait()
}
