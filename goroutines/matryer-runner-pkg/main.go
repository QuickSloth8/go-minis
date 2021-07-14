package main

import (
	"log"
	"math/rand"
	"time"

	runner "github.com/matryer/runner"
)

func randSleep() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n := r.Intn(10)
	time.Sleep(100 * time.Duration(n) * time.Millisecond)
}

func main() {
	task := runner.Go(func(shouldStop runner.S) error {
		defer func() {
			// tear-down logic
		}()

		for {
			log.Println("Running, still :)")
			randSleep()

			if shouldStop() {
				break
			}
		}

		return nil
	})

	time.Sleep(5 * time.Second)
	log.Println("requesting stoppage")
	task.Stop()
	select {
	case <-task.StopChan():
		log.Println("task successfully stopped")
	case <-time.After(1 * time.Second):
		log.Println("task stoppage time exceeded timeout!")
	}

	log.Println("MAIN END")
}
