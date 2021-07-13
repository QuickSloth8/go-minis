package main

import (
	"errors"
	"log"
	"strconv"
)

// NOTE: if the function had the following definition:
// GenerateNumbersRange(start int, end int) (numbers_chan chan int, err error)
// Then there will be a bug, and the programs would deadlock.
// The reason is, declaring a named return: numbers_chan
// (still don't have the exact explanation)
func GenerateNumbersRange(start int, end int) (chan int, error) {
	if start >= end {
		err := errors.New("start should be smaller than end")
		return nil, err
	}

	numbers_chan := make(chan int)

	go func() {
		defer close(numbers_chan)
		for i := start; i < end; i++ {
			log.Println(`pushing ` + strconv.Itoa(i))
			numbers_chan <- i
		}
	}()

	return numbers_chan, nil

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
