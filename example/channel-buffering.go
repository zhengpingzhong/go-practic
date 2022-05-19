package main

import "fmt"

func main() {
	messages := make(chan string, 8)
	messages <- "buffered"
	messages <- "channel"
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
