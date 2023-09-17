package main

import "fmt"

func main() {

	// Allows goroutines to communicate
	// as memory access is shared
	// The below code blocks since its unbuffered channel and send & recieve are blocked
	// on each other.
	// c := make(chan int)
	// c <- 2
	// <-c

	// Notice 1 is capacity not length
	// length is number of elemts inside channel
	// Also below is buffered channel and send & recieve are non-blocking
	// Send only blocks when channel is full
	// Recieve blocks when channel is empty
	c := make(chan int, 1)
	fmt.Println(len(c), cap(c))

	c <- 2
	fmt.Println(len(c))
	<-c

	// Since buffered channel blocks on recieve when channel is empty
	// and if you're sending multiple values in channel, its good idea to close it
	// to let the receibver no it can non-block
	c1 := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)
	}()

	// Equivalent to if val, ok := <- c1; ok
	// ok will be false when c1 is closed
	for val := range c1 {
		fmt.Println(val)
	}

	// Buffered channels will be useful when you want to compute something expensive without waiting on receiver
	// and buffer all of them in channel so that receiver
	// can easily pick them up once receiver is ready

	// select is just like switch but not top down
	// it waits until one of channel is non-blocking, if all of them blocking it selects
	// default case
	// If mutiple channels are recieved at same time, it selects one at random
	select {
	case <-c1:
	case <-c:
	default:
	}
}
