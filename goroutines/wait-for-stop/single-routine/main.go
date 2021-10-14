/*
A dedicated channel (stopChan) for signaling stop, and
another channel (stoppedChan) to wait for the goroutine stoppage.

P.S.: for multiple goroutines, it would be better to use
sync.WaitGroup to wait for all goroutines stoppage, as demonstrated in
the other project (under same directory).
*/

package main

import (
	"log"
	"time"
)

func main() {
	// a channel to tell it to stop
	stopChan := make(chan struct{})

	// a channel to signal that it's stopped
	stoppedChan := make(chan struct{})

	go func() { // work in background
		// close the stoppedChan when this func exits
		defer close(stoppedChan)

		// do setup work

		defer func() {
			// do teardown work
		}()

		for {
			select {
			default:
				log.Println("executing default case")
				time.Sleep(100 * time.Millisecond)
			case <-stopChan:
				// stop
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)

	log.Println("stopping...")

	close(stopChan) // tell it to stop (no need to send a value)
	<-stoppedChan   // wait for it to stop

	log.Println("Stopped.")
}
