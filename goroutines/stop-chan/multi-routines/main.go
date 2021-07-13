/*
stopping goroutines with a dedicated chan, and waiting for
graceful termination with sync.WaitGroup.

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

func foo(i int, stopchan chan bool) {
	for {
		select {
		default:
			log.Println(`I'm routine no. ` + strconv.Itoa(i))
			time.Sleep(RandomDuration())
		case <-stopchan:
			// stop logic
			return
		}
	}
}

func wrapper(wg *sync.WaitGroup, num_of_goroutines int, stop_after time.Duration) {
	stopchan := make(chan bool)

	for i := 1; i <= num_of_goroutines; i++ {
		wg.Add(1) // signal new foo instance
		go func(i int) {
			defer wg.Done() // signal end of foo
			foo(i, stopchan)
		}(i)
	}
	time.Sleep(stop_after) // request goroutines termination after ...
	close(stopchan)
}

func main() {
	var wg sync.WaitGroup

	num_of_goroutines := 5
	stop_after := 5 * time.Second

	wrapper(&wg, num_of_goroutines, stop_after)

	wg.Wait()
	log.Println("DONE!")
}
