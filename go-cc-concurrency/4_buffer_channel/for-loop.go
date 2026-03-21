package main

import "fmt"

func main() {
	charChannel := make(chan string, 5) // buffer channel

	chars := []string{"a", "b", "c", "d", "e"}

	for _, s := range chars {
		select { // but why select statement used?
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

}
