package b

type Writer interface {
	Write(p []byte) (n int, err error)
}
