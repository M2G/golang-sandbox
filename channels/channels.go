package channels

import (
	"fmt"
	"time"
)

func Channel() {
	message := make(chan string)

	go func() {
		fmt.Println("B: Sending message to channel")
		message <- "ping"
		fmt.Println("B: Message sent to channel")
	}()

	fmt.Println("A: Doing some work before receiving message from channel")
	time.Sleep(500 * time.Millisecond)

	msg := <-message

	fmt.Println("A: Received message from channel")
	fmt.Println(msg)
	time.Sleep(100 * time.Millisecond)
}
