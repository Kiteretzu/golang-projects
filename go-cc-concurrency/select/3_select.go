package main

import "fmt"

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		anotherChannel <- "cowy"
	}()
	go func() {
		myChannel <- "data"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnother := <-anotherChannel:
		fmt.Println(msgFromAnother)
	}

}
