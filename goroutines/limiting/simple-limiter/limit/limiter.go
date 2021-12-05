package limit

import "errors"

type limiter struct {
	c   chan struct{}
	log func(string)
}

// NewLimiter creates a new limiter with maxGoroutinesNum goroutines.
func NewLimiter(maxGoroutinesNum int, log func(message string)) *limiter {
	return &limiter{
		c:   make(chan struct{}, maxGoroutinesNum),
		log: log,
	}
}

// ExecBegin notifies the limiter about the start of work.
func (limiter *limiter) ExecBegin() {
	if limiter.log != nil {
		limiter.log("ExecBegin - start")
		defer limiter.log("ExecBegin - end")
	}

	limiter.c <- struct{}{}
}

// ExecEnd notifies the limiter about the end of work.
func (limiter *limiter) ExecEnd() error {
	if limiter.log != nil {
		limiter.log("ExecEnd - start")
		defer limiter.log("ExecEnd - end")
	}

	if len(limiter.c) <= 0 {
		return errors.New("there are no running goroutines")
	}
	<-limiter.c

	return nil
}
