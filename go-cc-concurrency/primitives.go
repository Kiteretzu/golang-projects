package main

import (
	"fmt"
	"time"
)

func comeFunc(num string) {
	fmt.Println(num)
}

func main () {

	go comeFunc("1")
	go comeFunc("2")
	go comeFunc("3")

 time.Sleep(time.Second * 2)

	fmt.Println("hi")
} 