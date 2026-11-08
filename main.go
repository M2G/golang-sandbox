package main

import (
	//"golang-sandbox/sync"
	"fmt"
	"golang-sandbox/channels"

	//"golang-sandbox/goroutine"
)

func main() {
	// sync.SyncPool()
	// sync.SyncWaitGroup()
	// sync.SyncMutex()
	// sync.Once()
	// sync.SyncOnceFunc()
	//goroutine.Say("Hello world from goroutine!")
	channels.Channel()

	fmt.Println("Hello world from main!")
}
