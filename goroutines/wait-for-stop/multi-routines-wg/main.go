/*
stopping multiple goroutines with a dedicated chan, and waiting
for graceful termination with sync.WaitGroup.

P.S.: added some extra features like randomness and timeouts
to imitate the behavior of web services
*/

package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func RandomDuration() time.Duration {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	i := r.Intn(10)
	return 100 * time.Duration(i) * time.Millisecond

}

func foo(i int, stopChan chan bool) {
	for {
		select {
		default:
			log.Println(`I'm routine no. ` + strconv.Itoa(i))
			time.Sleep(RandomDuration())
		case <-stopChan:
			// stop logic
			return
		}
	}
}

func wrapper(wg *sync.WaitGroup, num_of_goroutines int, stop_after time.Duration) {
	stopChan := make(chan bool)

	for i := 1; i <= num_of_goroutines; i++ {
		wg.Add(1)                // signal new foo instance
		go func(routineId int) { // important note: passing loop variables are required, for value preservation
			defer wg.Done() // signal end of foo
			foo(routineId, stopChan)
		}(i)
	}
	time.Sleep(stop_after) // request goroutines termination after ...
	close(stopChan)
}

func main() {
	var wg sync.WaitGroup

	numOfGoroutines := 5
	stopAfter := 5 * time.Second

	wrapper(&wg, numOfGoroutines, stopAfter)

	wg.Wait()
	log.Println("DONE!")
}
