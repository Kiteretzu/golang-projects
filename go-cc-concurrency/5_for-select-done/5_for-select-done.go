package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Done")
			return
		default:
			fmt.Println("Doing work")
		}
	}
}

func main() {

	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 3)

	close(done)

}
