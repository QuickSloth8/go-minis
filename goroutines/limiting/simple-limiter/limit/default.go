package limit

import "runtime"

var (
	TotalCoresNum  int
	defaultLimiter *limiter
)

func init() {
	TotalCoresNum = runtime.NumCPU()
	defaultLimiter = NewLimiter(TotalCoresNum, nil)
}

func ExecBeginDefault() {
	defaultLimiter.ExecBegin()
}

func ExecEndDefault() error {
	return defaultLimiter.ExecEnd()
}
