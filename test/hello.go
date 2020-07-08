package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutin(s string, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go goroutin("World", &wg)
	normal("Hello")
	// time.Sleep(2000 * time.Microsecond)
	wg.Wait()
}
