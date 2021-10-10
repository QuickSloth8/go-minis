/*
Demonstrates how to use a channel to send values,
and stop when there are no more values (when chan is stopped).

The channel is initialized and stopped the same
function that starts the goroutine.
*/
package main

import (
	"errors"
	"log"
	"strconv"
)

// NOTE: if the function had the following definition:
// GenerateNumbersRange(start int, end int) (numbers_chan chan int, err error),
// then there will be a bug, and the program would deadlock.
// The reason is, declaring a named return: numbers_chan
// (still don't have the exact explanation).

func GenerateNumbersRange(start int, end int) (chan int, error) {
	if start >= end {
		err := errors.New("start should be smaller than end")
		return nil, err
	}

	numbersChan := make(chan int)

	go func() {
		defer close(numbersChan)
		for i := start; i < end; i++ {
			log.Println(`pushing `, strconv.Itoa(i))
			numbersChan <- i
		}
	}()

	return numbersChan, nil

}

func main() {
	c, err := GenerateNumbersRange(1, 10)

	if err != nil {
		log.Fatal(err)
	}

	for n := range c {
		log.Println(`recieved ` + strconv.Itoa(n))
	}

	log.Println("MAIN END")
}
