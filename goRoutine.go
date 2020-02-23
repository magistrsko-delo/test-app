package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{} // sync wait group.. sysnc go routines
var counter = 0
var m = sync.RWMutex{} // read write mutex for go roteen sync... If anything is reading we cant write to it at all

func main()  {
	// abstraction of thread... map go routine on thread...
	// no interection og low level thread
	// reallocated go routines..
	// go sayHello() // green thread, thread pooling
	// say hello don't have time
	/*var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg) // passing message by value... copy of message
	msg = "Goodbye"*/

	fmt.Println(runtime.GOMAXPROCS(-1))
	fmt.Println()
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}

	wg.Wait()
}

func sayHello()  {
	// m.RLock() // readLock
	fmt.Println(counter)
	m.RUnlock()
	wg.Done()
}

func increment()  {
	// m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}