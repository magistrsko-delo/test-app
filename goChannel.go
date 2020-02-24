package main

import (
	"fmt"
	"sync"
	time2 "time"
)

// data between goRoutines
var wg1 = sync.WaitGroup{}

const (
	logInfo = "INFO"
	logWarning = "WARNING"
	logError = "ERROR"
)

type logEntry struct {
	time time2.Time
	severity string
	message string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel

func main()  {
	// ch := make(chan int, 50) // channel.. strongly type.. sender or receiver need a time for process
	// wg1.Add(2)

	/*go func(ch <- chan int) { // receive only channel
		i := <- ch // receiving go routinee, pulling data from the channel
		fmt.Println(i)
		i = <- ch // receiving go routinee, pulling data from the channel
		fmt.Println(i)

		for i := range ch {
			fmt.Println(i)
		}

		for {
			i, ok := <- ch
			if ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg1.Done()
	}(ch)

	go func(ch chan <- int) { // send only channel
		i := 42
		ch <- i // sending goRoutine, push data into go routine when there is space

		i = 27
		ch <- i
		close(ch) // closing channel, defer for log channel
		wg1.Done()
	}(ch)

	for j := 0; j < 5; j++ {

	}*/

	go logger()
	logCh <- logEntry{
		time:     time2.Now(),
		severity: logInfo,
		message:  "App is starting",
	}

	logCh <- logEntry{
		time:     time2.Now(),
		severity: logInfo,
		message:  "App is shutting down",
	}
	time2.Sleep(1000 * time2.Millisecond)
	doneCh <- struct{}{}
	// wg1.Wait()
}

// struct{} -- zero memmory allocation
// blocking select channel for logging

func logger()  {
	for {
		select {
		case entry := <- logCh:
			fmt.Println(  entry.time.String() + "[" + entry.severity + "] " + entry.message)
		case <- doneCh:
			break
		}

	}
}
