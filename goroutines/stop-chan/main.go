/*
A dedicated channel (stopchan) for signaling stop, and
another channel (stoppedchan) to wait for the goroutine stoppage.

P.S.: for multiple goroutines, it would be better to use
sync.WaitGroup to wait for all goroutines stoppage.
*/

package main

import (
	"log"
	"time"
)

func main() {
	// a channel to tell it to stop
	stopchan := make(chan struct{})

	// a channel to signal that it's stopped
	stoppedchan := make(chan struct{})

	go func() { // work in background
		// close the stoppedchan when this func exits
		defer close(stoppedchan)

		// do setup work

		defer func() {
			// do teardown work
		}()

		for {
			select {
			default:
				log.Println("executing default case")
				time.Sleep(100 * time.Millisecond)
			case <-stopchan:
				// stop
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)

	log.Println("stopping...")

	close(stopchan) // tell it to stop (no need to send a value)
	<-stoppedchan   // wait for it to stop

	log.Println("Stopped.")
}
